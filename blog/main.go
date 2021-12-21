package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	tc "gunkBlog/blog/core/category"
	tp "gunkBlog/blog/core/post"
	"gunkBlog/blog/services/category"
	"gunkBlog/blog/services/post"
	"gunkBlog/blog/storage/postgres"
	tcb "gunkBlog/gunk/v1/category"
	tpb "gunkBlog/gunk/v1/post"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("blog/env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Printf("error loading configuration: %v", err)
	}

	grpcServer := grpc.NewServer()
	store, err := newDBFromConfig(config)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	postStore, err := newDBFromConfig(config)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	cs := tc.NewCoreSvc(store)
	s := category.NewCategoryServer(cs)
	
	ps := tp.NewCoreSvc(postStore)
	sp := post.NewPostServer(ps)
	

	tcb.RegisterCategoryServiceServer(grpcServer, s)
	tpb.RegisterPostServiceServer(grpcServer, sp)

	host, port := config.GetString("server.host"), config.GetString("server.port")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	log.Printf("Server is starting on: http://%s:%s", host, port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func newDBFromConfig(config *viper.Viper) (*postgres.Storage, error) {
	cf := func(c string) string { return config.GetString("database." + c) }
	ci := func(c string) string { return strconv.Itoa(config.GetInt("database." + c)) }
	dbParams := " " + "user=" + cf("user")
	dbParams += " " + "host=" + cf("host")
	dbParams += " " + "port=" + cf("port")
	dbParams += " " + "dbname=" + cf("dbname")
	if password := cf("password"); password != "" {
		dbParams += " " + "password=" + password
	}
	dbParams += " " + "sslmode=" + cf("sslMode")
	dbParams += " " + "connect_timeout=" + ci("connectionTimeout")
	dbParams += " " + "statement_timeout=" + ci("statementTimeout")
	dbParams += " " + "idle_in_transaction_session_timeout=" + ci("idleTransacionTimeout")
	db, err := postgres.NewStorage(dbParams)
	if err != nil {
		return nil, err
	}
	return db, nil
}