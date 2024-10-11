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
		Id:   word.Id,
		Text: word.Text,
	}}, nil
}

func (s *WordService) UpdateWord(ctx context.Context, req *pb.UpdateWordRequest) (*pb.UpdateWordReply, error) {
	return &pb.UpdateWordReply{}, nil
}

func (s *WordService) DeleteWord(ctx context.Context, req *pb.DeleteWordRequest) (*pb.DeleteWordReply, error) {
	return &pb.DeleteWordReply{}, s.uc.DeleteWord(ctx, req.Id)
}

func (s *WordService) GetWord(ctx context.Context, req *pb.GetWordRequest) (*pb.GetWordReply, error) {
	return &pb.GetWordReply{}, nil
}

func (s *WordService) ListWord(ctx context.Context, req *pb.ListWordRequest) (*pb.ListWordReply, error) {
	words, err := s.uc.ListWords(ctx, req.Group)
	if err != nil {
		return nil, err
	}
	result := make([]*pb.Word, 0)
	for _, word := range words {
		result = append(result, &pb.Word{
			Id:   word.Id,
			Text: word.Text,
		})
	}
	return &pb.ListWordReply{
		Words: result,
	}, nil
}
