package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sparklee/abc-api/internal/biz"
)

type wordRepo struct {
	data *Data
	log  *log.Helper
}

func NewWordRepo(data *Data, logger log.Logger) biz.WordRepo {
	return &wordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *wordRepo) Save(ctx context.Context, word *biz.Word) (*biz.Word, error) {
	_, err := r.data.db.Words.
		Create().
		SetWord(word.Text).
		SetGroup("foo").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (r *wordRepo) List(ctx context.Context, group string) ([]*biz.Word, error) {
	words, err := r.data.db.Words.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*biz.Word, 0)
	for _, word := range words {
		result = append(result, &biz.Word{
			Text: word.Word,
		})
	}
	return result, nil
}
