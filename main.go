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
	fmt.Print("Hola busca 0.1\n")
	db, _ := sql.Open("mysql", "root:123456@/")
	defer db.Close()

	// Connect and check the server version
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":5000", nil)
}