package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const DRIVERNAME = "mysql"
const USER = "root"
const PASS = "vitalik88"
const NAMEDB = "tcp(localhost:3306)/usersdb"

type User struct {
	Id             uint16 `json:"id"`
	Email          string `json:"email"`
	ActivationCode int    `json:"activation_code"`
	Status         string `json:"status"`
	Description    string `json:"description"`
}

func (u User) Insert(Email string, Code int) {
	db, err := sql.Open(DRIVERNAME, USER+":"+PASS+"@"+NAMEDB)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `usersdb`.`Users` (`email`, `activation_code`) VALUES (?, ?)", Email, Code)
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}

func (u User) GetEmail(Email string) string {
	db, err := sql.Open(DRIVERNAME, USER+":"+PASS+"@"+NAMEDB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("select email from usersdb.Users where email = ?", Email)
	for res.Next() {
		var user User
		err = res.Scan(&user.Email)
		if err != nil {
			panic(err.Error())
		}
		email := user.Email
		return email
	}
	return ""
}
