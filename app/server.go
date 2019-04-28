package app

import (
	"context"
	"github.com/soichisumi/grpc-template/app/proto"
)

func NewServer() Server {return Server{}}

type Server struct{}

func (s Server) GetData(ctx context.Context, req *proto.GetDataRequest) (*proto.GetDataResponse, error) {
	return &proto.GetDataResponse{
		Str: "result",
		Num: 5,
	}, nil
}
