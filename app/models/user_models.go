package models

import (
	"belajar-golang/connection"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Username       string
	Email          string
	Password       string
}

func GetAllUsers() ([]User, error) {
	rows, err := connection.DB.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		fmt.Println("Error querying database:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error after scanning rows:", err.Error())
		return nil, err
	}
	return users, nil
}

func CreateUser(username, email, password string) error {
	_, err := connection.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)
	if err != nil {
		fmt.Println("Error creating user:", err.Error())
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	_, err := connection.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		fmt.Println("Error deleting user:", err.Error())
		return err
	}
	return nil
}

func AuthenticateUser(username, password string) (*User, error) {
	var user User
	err := connection.DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("Error querying database:", err.Error())
		return nil, errors.New("invalid credentials")
	}
	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Error verifying password:", err.Error())
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}
