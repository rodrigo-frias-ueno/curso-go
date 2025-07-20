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

// @title Hendyla Categories Api
// @version 1.0
// @description This is an api to manage item categories.
// @termsOfService http://swagger.io/terms/

// @contact.name   Hendyla Team
// @contact.url    http://www.swagger.io/support
// @contact.email  hendyla@itti.digital

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host hendyla.com
// @BasePath /api/v1
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
