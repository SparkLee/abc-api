package service

import (
	"context"

	pb "github.com/sparklee/abc-api/api/abc/v1"
)

type WordServiceService struct {
	pb.UnimplementedWordServiceServer
}

func NewWordServiceService() *WordServiceService {
	return &WordServiceService{}
}

func (s *WordServiceService) CreateWord(ctx context.Context, req *pb.CreateWordRequest) (*pb.CreateWordReply, error) {
	return &pb.CreateWordReply{}, nil
}
func (s *WordServiceService) UpdateWord(ctx context.Context, req *pb.UpdateWordRequest) (*pb.UpdateWordReply, error) {
	return &pb.UpdateWordReply{}, nil
}
func (s *WordServiceService) DeleteWord(ctx context.Context, req *pb.DeleteWordRequest) (*pb.DeleteWordReply, error) {
	return &pb.DeleteWordReply{}, nil
}
func (s *WordServiceService) GetWord(ctx context.Context, req *pb.GetWordRequest) (*pb.GetWordReply, error) {
	return &pb.GetWordReply{}, nil
}
func (s *WordServiceService) ListWord(ctx context.Context, req *pb.ListWordRequest) (*pb.ListWordReply, error) {
	return &pb.ListWordReply{}, nil
}
