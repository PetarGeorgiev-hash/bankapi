package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/PetarGeorgiev-hash/bankapi/api"
	db "github.com/PetarGeorgiev-hash/bankapi/db/sqlc"
	"github.com/PetarGeorgiev-hash/bankapi/gapi"
	"github.com/PetarGeorgiev-hash/bankapi/pb"
	_ "github.com/PetarGeorgiev-hash/bankapi/swagger/statik"
	"github.com/PetarGeorgiev-hash/bankapi/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	var err error
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't read env file", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	runDbMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Can't create server", err.Error())
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Can't start server", err.Error())
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Can't create server", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBankAPIServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Can't create listener", err.Error())

	}
	log.Printf("Starting gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Can't start grpc server", err.Error())

	}
}

func runGatewayServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Can't create server", err.Error())
	}

	jsonOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOptions)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterBankAPIServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("Can't register server", err.Error())
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("Can't create statik fs", err.Error())
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Can't create listener", err.Error())

	}
	log.Printf("Starting HTTP Gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Can't start HTTP Gateway server", err.Error())

	}
}

func runDbMigration(url string, dbSource string) {
	migration, err := migrate.New(url, dbSource)
	if err != nil {
		log.Fatal("Can't create new migrate instance ", err.Error())
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Can't run migrate up ", err.Error())
	}

	log.Println("DB migrated successfully")
}
