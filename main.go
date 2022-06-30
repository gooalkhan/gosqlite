package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       string
	UserId   string
	Password string
}

func InitDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	createTableQuery := `
		create table IF NOT EXISTS useraccount ( 
		id integer PRIMARY KEY autoincrement,
		userId text,
		password text,
		UNIQUE (id, userId)
		)
	`
	_, e := db.Exec(createTableQuery)
	if e != nil {
		return nil, e
	}

	return db, nil
}

func AddUser(db *sql.DB, id string, password string) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into useraccount (userId,password) values (?,?)")
	_, err := stmt.Exec(id, password)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func GetUser(db *sql.DB, userId string) (User, error) {
	var user User
	rows := db.QueryRow("select * from useraccount where userId = $1", userId)
	err := rows.Scan(&user.Id, &user.UserId, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func main() {
	fmt.Println("Hello, World!")
	db, err := InitDB("file::memory:?mode=memory&cache=shared")
	if err != nil {
		log.Println(err)
	}
	log.Println("DB initialized")
	e := AddUser(db, "jinyoung", "1234")
	if e != nil {
		log.Println(err)
	}
	log.Println("Record added")
	u, er := GetUser(db, "jinyoung")
	if er != nil {
		log.Println(er)
	}
	fmt.Println(u)
	fmt.Println("program end")

}
