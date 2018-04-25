package main

import (
	"net"
	"net/http"
	"io"
	"log"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	pb "eating.com/app/auth"
)

type User struct {
    Id          int
    LoginName   string
    Password	string
    Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
    Id          int
    Name   		string
    Age         int16
}

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
   	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(mariadb:3306)/eating?charset=utf8")
	orm.RegisterModel(new(User), new(Profile))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

var version string

type server struct { }

func (s *server) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	err := o.Begin()

    if err != nil {
    	return nil, err
    }
    
    profile := new(Profile)	
    profile.Name = r.Username
    profile.Age = 30

    user := new(User)
    user.Profile = profile
    user.LoginName = r.Username
    user.Password = r.Password

	_, err = o.Insert(profile)
	if err != nil {
		o.Rollback()
		return nil, err
	} 

	userId, err := o.Insert(user)
	if err != nil {
		o.Rollback()
		return nil, err
	} 

	err = o.Commit()
	if err != nil {
		o.Rollback()
		return nil, err
	} 


	rr := new(pb.LoginResponse)
	rr.Token = fmt.Sprint("HAHAHddA", userId)
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

/*
CREATE DATABASE eating
  DEFAULT CHARACTER SET utf8
  DEFAULT COLLATE utf8_general_ci;

SELECT *
FROM user u INNER JOIN profile p ON u.profile_id = p.id
*/
var (
	o orm.Ormer
)

func main() {
	lis, errr := net.Listen("tcp", ":5000")
	if errr != nil {
		log.Fatalf("failed to listen: %v", errr)
	}

	o = orm.NewOrm()
    o.Using("default") 

	err := orm.RunSyncdb("default", true, true)
	if err != nil {
	    log.Fatalf(err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer,  &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

