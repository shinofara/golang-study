package main_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	baseURL = "http://localhost:8888"
	userNum = 3
)

type Transaction struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func TestCreate(t *testing.T) {
	limitMap := make(map[int]int, userNum)
	for uID := 1; uID <= userNum; uID++ {
		limitMap[uID] = uID * 1000
	}

	// Create transactions
	var wg sync.WaitGroup
	wg.Add(userNum)

	for uID := 1; uID <= userNum; uID++ {
		go func(uID int) {
			defer wg.Done()
			var total int
			for j := 0; j <20; j++ {
				buffer := bytes.NewBuffer(make([]byte, 0, 128))
				amount := uID * 100
				total += amount
				if err := json.NewEncoder(buffer).Encode(Transaction{
					UserID:      uID,
					Amount:      amount,
					Description: fmt.Sprintf("商品%d", uID),
				}); err != nil {
					t.Fatal(err)
				}
				req, err := http.NewRequest(
					http.MethodPost,
					baseURL+"/transactions",
					buffer,
				)
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("apikey", fmt.Sprintf("secure-api-key-%d", uID))

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Fatal(err)
				}

				limit := limitMap[uID]
				if total > limit {
					want := http.StatusPaymentRequired
					if resp.StatusCode != want {
						t.Errorf("POST /transactions status %d != %d total:%d limit:%d", resp.StatusCode, want, total, limit)
					}
				} else {
					want := http.StatusCreated
					if resp.StatusCode != want {
						t.Errorf("POST /transactions status %d != %d total:%d limit:%d", resp.StatusCode, want, total, limit)
					}
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatal(err)
				}
				t.Log(string(body))

				if err := resp.Body.Close(); err != nil {
					t.Fatal(err)
				}
			}
		}(uID)
	}
	wg.Wait()

	// Check limit
	conn, err := sql.Open("mysql", "root@tcp(127.0.0.1:43306)/codetest")
	if err != nil {
		t.Fatal(err)
	}
	for uID := 1; uID <= userNum; uID++ {
		var amount int
		if err := conn.QueryRow(
			"select sum(amount) from transactions where user_id=?",
			uID,
		).Scan(&amount); err != nil {
			t.Fatal(err)
		}
		limit := limitMap[uID]
		if amount > limit {
			t.Errorf("amount %d over the limit %d", amount, limit)
		}
	}
}
