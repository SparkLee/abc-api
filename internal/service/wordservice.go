package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/sparklee/abc-api/api/abc/v1"
	"github.com/sparklee/abc-api/internal/biz"
)

var AliyunClient *sdk.Client

type WordService struct {
	pb.UnimplementedWordServiceServer
	log *log.Helper
	uc  *biz.WordUsecase
}

type TokenResult struct {
	ErrMsg string
	Token  struct {
		UserId     string
		Id         string
		ExpireTime int64
	}
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

func (s *WordService) GetAliyunNlsToken(ctx context.Context, req *pb.GetAliyunNlsTokenRequest) (*pb.GetAliyunNlsTokenReply, error) {
	client := AliyunClient
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Domain = "nls-meta.cn-shanghai.aliyuncs.com"
	request.ApiName = "CreateToken"
	request.Version = "2019-02-28"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpStatus())
	fmt.Print(response.GetHttpContentString())

	var tr TokenResult
	err = json.Unmarshal([]byte(response.GetHttpContentString()), &tr)
	if err == nil {
		fmt.Println(tr.Token.Id)
		fmt.Println(tr.Token.ExpireTime)
	} else {
		fmt.Println(err)
	}

	return &pb.GetAliyunNlsTokenReply{Token: tr.Token.Id}, nil
}
