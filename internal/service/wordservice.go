package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sparklee/abc-api/internal/biz"

	pb "github.com/sparklee/abc-api/api/abc/v1"
)

type WordService struct {
	pb.UnimplementedWordServiceServer
	log *log.Helper
	uc  *biz.WordUsecase
}

func NewWordService(uc *biz.WordUsecase, logger log.Logger) *WordService {
	return &WordService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *WordService) CreateWord(ctx context.Context, req *pb.CreateWordRequest) (*pb.CreateWordReply, error) {
	s.log.Infof("input data: %v", req)
	word, err := s.uc.CreateWord(ctx, &biz.Word{
		Text: req.Text,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateWordReply{Word: &pb.Word{
		Id:   1,
		Text: word.Text,
	}}, nil
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
