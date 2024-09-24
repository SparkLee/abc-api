package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/sparklee/abc-api/api/abc/v1"
)

type WordService struct {
	pb.UnimplementedWordServiceServer
	log *log.Helper
}

func NewWordService() *WordService {
	return &WordService{}
}

func (s *WordService) CreateWord(ctx context.Context, req *pb.CreateWordRequest) (*pb.CreateWordReply, error) {
	return &pb.CreateWordReply{}, nil
}
func (s *WordService) UpdateWord(ctx context.Context, req *pb.UpdateWordRequest) (*pb.UpdateWordReply, error) {
	return &pb.UpdateWordReply{}, nil
}
func (s *WordService) DeleteWord(ctx context.Context, req *pb.DeleteWordRequest) (*pb.DeleteWordReply, error) {
	return &pb.DeleteWordReply{}, nil
}
func (s *WordService) GetWord(ctx context.Context, req *pb.GetWordRequest) (*pb.GetWordReply, error) {
	return &pb.GetWordReply{}, nil
}
func (s *WordService) ListWord(ctx context.Context, req *pb.ListWordRequest) (*pb.ListWordReply, error) {
	return &pb.ListWordReply{}, nil
}
