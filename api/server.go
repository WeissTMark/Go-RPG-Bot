package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var calls = map[string]func(w http.ResponseWriter, r *http.Request){
	"/":      hello,
	"/login": login,
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bot")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {

	for route, fun := range calls {
		http.HandleFunc(route, fun)
	}

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello world!"}`))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	db := dbConn()

	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")
	fmt.Println("INSERT: User: " + user + " | Pass: " + pass)
	insForm, err := db.Prepare("INSERT INTO users(username, pass) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(user, pass)

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
