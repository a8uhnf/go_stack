package main

import (
	"log"
	"net"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"sample-project/proxy"
	pb "sample-project/proto"

	"fmt"
	"time"
)

const (
	port = ":50051"
	pg_username = "vagrant"
	pg_password = "vagrant"
	pg_port = "5432"
	pg_dbname = "test_xorm"
	pg_host = "localhost"
)

type Person struct {
	Id int64
	Name string `xorm:"varchar(25)"`
	Responsibility string `xorm:"varchar(25)"`
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("Checkout..", in.Name)

	eng, err := NewEngine(pg_username, pg_password, pg_host, pg_port, pg_dbname)
	if err != nil {
		return nil, err
	}
	log.Println("DB Connected...")
	nm := new(Person)
	nm.Name = in.Name
	nm.Responsibility = ""
	affected, err := eng.Insert(nm)
	log.Println("------------", affected)
	if err != nil {
		return nil, err
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func (s *server) SayHelloEmpty(ctx context.Context, in *pb.VoidRequest) (*pb.HelloReply, error) {
	log.Println("Checkout..")
	return &pb.HelloReply{Message: "Hello From Empty World"}, nil
}


func NewEngine(username, password, host, port, dbName string) (*xorm.Engine, error) {
	conn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		username, password, host, port, dbName)
	engine, err := xorm.NewEngine("postgres", conn)
	if err != nil {
		return nil, err
	}
	engine.SetMaxIdleConns(0)
	engine.DB().SetConnMaxLifetime(10 * time.Minute)
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	return engine, err
}


func main() {
	log.Println("Server Started...")


	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Error Occured")
		return
	}

	s := grpc.NewServer()
	pb.RegisterSampleProjectServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go proxy.Initialize()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
