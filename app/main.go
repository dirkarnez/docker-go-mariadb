package main
import (
	"net/http"
	"io"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

var version string

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(mariadb:3306)/")
	defer db.Close()
	if err != nil {
			fmt.Print(err.Error())
	}
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":5000", nil)
}