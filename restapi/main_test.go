package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"reflect"
)

func Test_indexHandler(t *testing.T) {
	ts := httptest.NewServer( http.HandlerFunc( indexHandler ) )
	defer ts.Close()


	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		status int
	}{
		{
			"return status is 200",
			args{},
			200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := http.Get( ts.URL )
	        if err != nil {
                t.Error("unexpected")
	        }

	        if res.StatusCode != tt.status {
                t.Error("Status code error")
	        }

	        b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}

	        if string(b) != "1" {
	        	 t.Errorf("Response body error. body = %s", string(b))
	        }

	        if err := res.Body.Close(); err != nil {
		        t.Fatal(err)
	        }
		})
	}
}


func Test_noContentHandler(t *testing.T) {
	ts := httptest.NewServer( http.HandlerFunc( noContentHandler ) )
	defer ts.Close()


	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		status int
	}{
		{
			"return status is 204",
			args{},
			204,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := http.Get( ts.URL )
			if err != nil {
				t.Error("unexpected")
				return
			}

			if res.StatusCode != tt.status {
				t.Errorf("Status code error. return code = %d", res.StatusCode)
				return
			}

	        if err := res.Body.Close(); err != nil {
		        t.Fatal(err)
	        }
		})
	}
}

func Test_jsonHandler(t *testing.T) {
	ts := httptest.NewServer( http.HandlerFunc( jsonHandler ) )
	defer ts.Close()


	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		status int
		want []byte
	}{
		{
			"期待するJSONをレスポンスで返す",
			args{},
			200,
			[]byte(`{"year": "2019", "status": 200}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := http.Get( ts.URL )
	        if err != nil {
                t.Error("unexpected")
                return
	        }

	        if res.StatusCode != tt.status {
                t.Errorf("Status code error. return code = %d", res.StatusCode)
                return
	        }

            b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}

            if string(b) != string(tt.want) {
	        	 t.Errorf("Response body error. body = %s", string(b))
	        }

            if err := res.Body.Close(); err != nil {
		        t.Fatal(err)
	        }
		})
	}
}

func Test_fizzBuzzHandler(t *testing.T) {
	ts := httptest.NewServer( http.HandlerFunc( fizzBuzzHandler ) )
	defer ts.Close()


	type fizzbuzz struct {
		value string
	}


	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want fizzbuzz
	}{
		{
			"in 1 out 1",
			args{n: 1},
			fizzbuzz{value: "1"},
		},
		{
			"in 3 out Fizz",
			args{n: 3},
			fizzbuzz{value: "Fizz"},
		},
		{
			"in 5 out Buzz",
			args{n: 5},
			fizzbuzz{value: "Buzz"},
		},
		{
			"in 15 out FizzBuzz",
			args{n: 15},
			fizzbuzz{value: "FizzBuzz"},
		},
		{
			"in 75 out FizzBuzz",
			args{n: 75},
			fizzbuzz{value: "1FizzBuzz"},
		},
		{
			"in 99 out Fizz",
			args{n: 99},
			fizzbuzz{value: "Fizz"},
		},
		{
			"in 100 out Buzz",
			args{n: 100},
			fizzbuzz{value: "Buzz"},
		},
		{
			"in 101 out 101",
			args{n: 101},
			fizzbuzz{value: "101"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 values := url.Values{}
			values.Add("n", string(tt.args.n))
			res, err := http.Get( ts.URL + "?" + values.Encode())
	        if err != nil {
                t.Error("unexpected")
                return
	        }

            if err := res.Body.Close(); err != nil {
		        t.Fatal(err)
	        }

            b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}

			var bb fizzbuzz
			if err := json.Unmarshal(b, &bb); err != nil {
				t.Fatal(err)
			}

            if !reflect.DeepEqual(bb, tt.want) {
	        	 t.Errorf("Response body error. body = %v", bb)
	        }

            if err := res.Body.Close(); err != nil {
		        t.Fatal(err)
	        }
		})
	}
}

