package main

import (
	"encoding/json"
	"github.com/pignuante/go-todoList/server"
	"io/ioutil"
	"log"
)

type loginInfo struct {
	DbAddr string `json:"loginInfo"`
}

func main() {

	file, err := ioutil.ReadFile("./loginInfo/loginInfo.json")
	if err != nil {
		return
	}
	var info loginInfo
	json.Unmarshal(file, &info)

	if err := server.ListenAndServe(server.Config{
		Address:     ":8080",
		DatabaseURL: info.DbAddr,
	}); err != nil {
		log.Fatalln(err)
	}
}
