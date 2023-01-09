package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func main() {
	conf := loadConfig("config.yml")
	db = connectDb(conf)
	defer db.Close()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on",l.Addr())
	
	router := httprouter.New()
	router.GET("/schemas", ListSchemas)
	router.GET("/schemas/:schema/tables", ListTables)

    log.Fatal(http.Serve(l, router))
}

func ListSchemas(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemas := getSchemas(db)
	collection := map[string][]string {
		"schemas": schemas,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(collection)
}

func ListTables(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tables := getTablesBySchema(db, ps.ByName("schema"))
	if(len(tables) == 0){
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Length", "0")
		return
	}
	collection := map[string][]string {
		"tables": tables,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(collection)
}