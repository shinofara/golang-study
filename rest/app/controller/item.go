package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shinofara/golang-study/rest/app/middleware"
)

// Transaction transaction entity.
type Transaction struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

// TransactionController controller.
type TransactionController struct {
	Base
}

func (t *TransactionController) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uID := r.Context().Value(middleware.UserIDKey).(int)
	rows, err := t.DB.Open().QueryContext(
		ctx,
		"select id, user_id, amount, description from transactions where user_id=?",
		uID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	list := make([]Transaction, 0)
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		list = append(list, t)
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TransactionController) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	uID := r.Context().Value(middleware.UserIDKey).(int)
	var transaction Transaction
	if err := t.DB.Open().QueryRowContext(
		ctx,
		"select id, user_id, amount, description from transactions where id=? and user_id=?",
		id,
		uID,
	).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Amount,
		&transaction.Description,
	); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TransactionController) Create(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err)
			return
		}
	}()

	ctx := r.Context()
	uID := r.Context().Value(middleware.UserIDKey).(int)
	result, err := t.DB.Open().ExecContext(
		ctx,
		"insert into transactions (user_id, amount, description) values (?,?,?)",
		uID,
		transaction.Amount,
		transaction.Description,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	transaction.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TransactionController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	uID := r.Context().Value(middleware.UserIDKey).(int)
	result, err := t.DB.Open().ExecContext(
		ctx,
		"delete from transactions where id=? and user_id=?",
		id,
		uID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if affected == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
