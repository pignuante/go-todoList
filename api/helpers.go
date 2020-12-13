package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pignuante/go-todoList/db"
	"io"
	"log"
	"net/http"
	"strconv"
)

func must(err error) {
	if err != nil {
		log.Println("internal error", err)
		panic(internalError)
	}
}

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	must(err)
	writeJSON(w, lists)
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	list, err := db.GetTodoList(listID)
	must(err)
	writeJSON(w, list)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	must(json.NewEncoder(w).Encode(v))
}

func parseJSON(r io.Reader, v interface{}) {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		log.Println("parsing json body", err)
		panic(malformedInputError)
	}
}

func parseIntParam(r *http.Request, key string) int {
	vars := mux.Vars(r)
	if v, ok := vars[key]; ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	panic(malformedInputError)
}
