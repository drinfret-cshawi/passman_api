package db

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId   int            `json:"userid"`
	UserName string         `json:"username"`
	FullName sql.NullString `json:"fullname"`
	Email    sql.NullString `json:"email"`
}

func GetUsers() ([]User, error) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := db.Query("SELECT id, username, fullname, email FROM users ORDER BY id")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var users []User

	for rows.Next() {
		//var id int
		//var userName string
		//var fullName sql.NullString
		//var email sql.NullString
		//
		//err = rows.Scan(&id, &userName, &fullName, &email)
		//if err != nil {
		//	fmt.Println(err)
		//	return nil, err
		//}
		//
		//users = append(users, User{UserId: id, Login: userName, FullName: fullName, Email: email})
		var user User
		err = rows.Scan(&user.UserId, &user.UserName, &user.FullName, &user.Email)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserById(id int) (User, error) {
	var user User

	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	row := db.QueryRow("SELECT id, username, fullname, email FROM users WHERE id = $1", id)

	//var userName string
	//var fullName sql.NullString
	//var email sql.NullString
	//
	//err = row.Scan(&id, &userName, &fullName, &email)

	err = row.Scan(&user.UserId, &user.UserName, &user.FullName, &user.Email)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

func AddUser(username string, fullname string, password string, email string) (int, error) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	insertString := "INSERT INTO users(username, fullname, password, email) VALUES($1, $2, $3, $4) returning id;"
	var lastInsertID int
	err = db.QueryRow(insertString, username, fullname, password, email).Scan(&lastInsertID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return lastInsertID, nil
}

func DeleteUser(userId int) (int, error) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	deleteString := "DELETE FROM users WHERE id = $1;"
	result, err := db.Exec(deleteString, userId)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	nRows, _ := result.RowsAffected()

	return int(nRows), nil
}

func UpdateUser(userId int, username string, fullname string, password string, email string) (int, error) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	updateString := "UPDATE users SET username = $2, fullname = $3, password = $4, email = $5 WHERE id = $1;"
	result, err := db.Exec(updateString, userId, username, fullname, password, email)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	nRows, _ := result.RowsAffected()

	return int(nRows), nil
}
