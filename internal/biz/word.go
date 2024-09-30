package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Word struct {
	Id    int64
	Group string
	Text  string
}

type WordRepo interface {
	Save(ctx context.Context, word *Word) (*Word, error)
	List(ctx context.Context, group string) ([]*Word, error)
	Delete(ctx context.Context, id int64) error
}

type WordUsecase struct {
	repo WordRepo
	log  *log.Helper
}

func NewWordUseCase(repo WordRepo, logger log.Logger) *WordUsecase {
	return &WordUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *WordUsecase) CreateWord(ctx context.Context, word *Word) (*Word, error) {
	uc.log.WithContext(ctx).Infof("CreateWord: %v", word)
	return uc.repo.Save(ctx, word)
}

func (uc *WordUsecase) ListWords(ctx context.Context, group string) ([]*Word, error) {
	return uc.repo.List(ctx, group)
}

func (uc *WordUsecase) DeleteWord(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
