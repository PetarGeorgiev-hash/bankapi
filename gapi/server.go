package gapi

import (
	"fmt"

	db "github.com/PetarGeorgiev-hash/bankapi/db/sqlc"
	"github.com/PetarGeorgiev-hash/bankapi/pb"
	"github.com/PetarGeorgiev-hash/bankapi/token"
	"github.com/PetarGeorgiev-hash/bankapi/util"
	"github.com/PetarGeorgiev-hash/bankapi/worker"
)

type Server struct {
	pb.UnimplementedBankAPIServiceServer
	config          util.Config
	store           db.Store
	token           token.Maker
	taskDistributor worker.TaskDistributor
}

// Creates new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can't create token maker %w", err)
	}

	server := &Server{store: store, token: tokenMaker, config: config, taskDistributor: taskDistributor}

	return server, nil
}
