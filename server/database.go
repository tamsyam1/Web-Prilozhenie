package main

import (
	"context"
	"fmt"
	_ "html/template"
	_ "log"
	_ "net/http"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Fullname string
	Password string
	Role     string
}


func VerifyUser(fullname, password string, conn *pgx.Conn) (*User, error) {
    var user User
    err := conn.QueryRow(context.Background(),
        "SELECT \"FIO\", password, role FROM users WHERE \"FIO\" = $1 AND password = $2",
        fullname, password).Scan(&user.Fullname, &user.Password, &user.Role)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, fmt.Errorf("неверные учетные данные")
        }
        return nil, fmt.Errorf("ошибка базы данных: %w", err)
    }
    return &user, nil
}


