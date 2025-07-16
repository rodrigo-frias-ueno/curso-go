package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	_ "github.com/go-sql-driver/mysql" // MySQL driver

	"github.com/jmoiron/sqlx"
	"github.com/ueno-tecnologia-org/go-core/pkg/web"
)

func main() {
	initLogger()

	strConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "12345", "localhost", 3306, "curso_go")
	db, err := sqlx.Connect("mysql", strConn)

	router := buildRouter(db)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Starting rest server", "address", listener.Addr().String())
	if err := web.Run(listener, web.DefaultTimeouts, router); err != nil {
		log.Fatal(err)
	}
}
