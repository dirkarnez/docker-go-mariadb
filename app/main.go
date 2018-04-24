package main
import (
	"net"
	"net/http"
	"io"
	"fmt"
	"log"
	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	pb "eating.com/app/auth"
)

type User struct {
    Id          int
    Name        string
    Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
    Id          int
    Age         int16
}

func init() {
    //orm.RegisterDriver("mysql", orm.DRMySQL)
   // orm.RegisterDataBase("default", "mysql", "root:123456@tcp(mariadb:3306)/eating?charset=utf8")
	//orm.RegisterModel(new(User), new(Profile))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

var version string

type server struct {

}

func (s *server) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("See u")
	rr := new(pb.LoginResponse)
	rr.Token = "HAHAHddA"
	rr.Functions = append(rr.Functions, new(pb.Function))
	return rr, nil
}

func (s *server) SayHello(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	/*md*/_, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid auth token required")
	}


	// jwtToken, ok := md["authorization"]
	// if !ok {
	// 	return nil, grpc.Errorf(codes.Unauthenticated, "valid auth token required")
	// }

	/*_, claims, err := s.validateJwtToken(jwtToken[0])
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "valid auth token required: %v", err)
	}*/

	return &pb.Response{Token: "Hello, !"}, nil
}

func main() {
	lis, errr := net.Listen("tcp", ":5000")
	if errr != nil {
		log.Fatalf("failed to listen: %v", errr)
	}
	/*o := orm.NewOrm()
    o.Using("default") 
    
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
	    fmt.Println(err)
	}

    profile := new(Profile)	
    profile.Age = 30

    user := new(User)
    user.Profile = profile
    user.Name = "12345陳"
	o.Insert(profile)
	o.Insert(user)*/
	fmt.Println("Hello World")
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer,  &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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

