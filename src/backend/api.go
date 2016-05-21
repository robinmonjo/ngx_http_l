package main

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

const (
	contextDBKey = "db"
	bucket       = "backend"
)

// api model
type entry struct {
	Host    string `json:"host"`
	Backend string `json:"backend"`
}

// routes definition
type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

var routes = []route{
	route{
		"GET",
		"/entries",
		index,
	},
	route{
		"GET",
		"/entries/{host}",
		show,
	},
	route{
		"PUT",
		"/entries/{host}",
		update,
	},
	route{
		"POST",
		"/entries",
		create,
	},
	route{
		"DELETE",
		"/entries/{host}",
		destroy,
	},
}

func startApi(port string, db *bolt.DB) error {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.method).Path(route.path).Handler(apiHandler(db, route.handler))
	}
	return http.ListenAndServe(":"+port, router)
}

// handlers
func apiHandler(db *bolt.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, contextDBKey, db)
		next.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	var entries = []*entry{}

	if err := getDB(r).View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil //bucket do not exists yet
		}
		return b.ForEach(func(k, v []byte) error {
			entries = append(entries, &entry{string(k), string(v)})
			return nil
		})
	}); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := json.NewEncoder(w).Encode(entries); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func show(w http.ResponseWriter, r *http.Request) {

}

func create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var e entry
	if err := decoder.Decode(&e); err != nil {
		http.Error(w, err.Error(), 422)
	}

	if err := getDB(r).Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		return b.Put([]byte(e.Host), []byte(e.Backend))
	}); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(e)
}

func update(w http.ResponseWriter, r *http.Request) {

}

func destroy(w http.ResponseWriter, r *http.Request) {

}

// helpers
func getDB(r *http.Request) *bolt.DB {
	return context.Get(r, contextDBKey).(*bolt.DB)
}
