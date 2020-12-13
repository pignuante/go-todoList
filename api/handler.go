package api

import (
	"encoding/json"
	"github.com/pignuante/go-todoList/db"
	"github.com/pignuante/go-todoList/todo"
	"log"
	"net/http"
)

//func getTodoLists(w http.ResponseWriter, r *http.Request) {
//	lists, err := db.GetTodoLists()
//	if err != nil {
//		panic(internalError)
//	}
//	json.NewEncoder(w).Encode(lists)
//}

func writeInternalError(w http.ResponseWriter, err error) {
	log.Println("internal error", err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorResponse{"Internal Error"})
}

func createTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	parseJSON(r.Body, &req)
	todoList, err := db.CreateTodoList(req.Name)
	must(err)
	writeJSON(w, todoList)
}
func modifyTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	var req todo.List
	parseJSON(r.Body, &req)
	must(db.RenameTodoList(listID, req.Name))
}

func deleteTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	must(db.DeleteTodoList(listID))
}

func createTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	var req todo.Item
	parseJSON(r.Body, &req)

	item, err := db.CreateTodoItem(listID, req.Text, req.Done)
	must(err)
	writeJSON(w, item)
}

func modifyTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	itemID := parseIntParam(r, "item_id")
	var req todo.Item
	parseJSON(r.Body, &req)
	must(db.ModifyTodoItem(listID, itemID, req.Text, req.Done))
}

func deleteTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	itemID := parseIntParam(r, "item_id")
	must(db.DeleteTodoItem(listID, itemID))
}
