package db

import (
	"fmt"
)

type Password struct {
	Id       int    `json:"id"`
	Site     string `json:"site"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func GetPasswordsForUser(userId int) ([]Password, error) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := db.Query("SELECT id, site, login, password FROM passwords p WHERE user_id = $1 ORDER BY id", userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var passwords []Password

	for rows.Next() {
		var id int
		var site string
		var username string
		var password string

		err = rows.Scan(&id, &site, &username, &password)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		passwords = append(passwords, Password{Id: id, Site: site, Login: username, Password: password})
	}

	return passwords, nil
}

func AddPassword(userId int, site string, username string, password string) (int, error) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	insertString := "INSERT INTO passwords(user_id, site, login, password) VALUES($1, $2, $3, $4) returning id;"

	var lastInsertID int
	err = db.QueryRow(insertString, userId, site, username, password).Scan(&lastInsertID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return lastInsertID, nil
}

func DeletePassword(userId int, site string, username string) (int, error) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	deleteString := "DELETE FROM passwords WHERE user_id = $1 AND site = $2 AND login = $3;"
	result, err := db.Exec(deleteString, userId, site, username)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	nRows, _ := result.RowsAffected()

	return int(nRows), nil
}

func UpdatePassword(userId int, site string, username string, password string) (int, error) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	updateString := "UPDATE passwords SET password = $4 WHERE user_id = $1 AND site = $2 AND login = $3;"
	result, err := db.Exec(updateString, userId, site, username, password)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	nRows, _ := result.RowsAffected()

	return int(nRows), nil
}
