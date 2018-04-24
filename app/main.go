package main
import (
	"net/http"
	"io"
	"fmt"
	//"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    Id          int
    Name        string
    Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
    Id          int
    Age         int16
    User        *User   `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
	orm.RegisterModel(new(User), new(Profile))
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:123456@tcp(mariadb:3306)/eating")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

var version string

func main() {
	o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database
    
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
	    fmt.Println(err)
	}

    profile := new(Profile)
    profile.Age = 30

    user := new(User)
    user.Profile = profile
    user.Name = "slene"

    fmt.Println(o.Insert(profile))
    fmt.Println(o.Insert(user))

	// db, err := sql.Open("mysql", "root:123456@tcp(mariadb:3306)/")
	// defer db.Close()
	// if err != nil {
	// 		fmt.Print(err.Error())
	// }
	// db.QueryRow("SELECT VERSION()").Scan(&version)
	// fmt.Println("Connected to:", version)
	// http.HandleFunc("/", helloHandler)
	// http.ListenAndServe(":5000", nil)
}

