package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func create(name string) {
    db, err := sql.Open("mysql", "root:mysqlgo@tcp(172.17.0.2:3306)/")
    if err != nil {
        panic(err)
    }
    fmt.Printf("1\n")
    defer db.Close()
    _,err = db.Exec("CREATE DATABASE " + name)
    if err != nil {
        panic(err)
    }
    fmt.Printf("2\n")
    _,err = db.Exec("USE " + name)
    if err != nil {
        panic(err)
    }
    fmt.Printf("3\n")
    _,err = db.Exec("CREATE TABLE users ( id integer, data varchar(32) )")
    if err != nil {
        panic(err)
    }
    fmt.Printf("4\n")
}

func main() {
    create("testdb")
}
