package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sparklee/abc-api/internal/biz"
	"github.com/sparklee/abc-api/internal/data/ent"
	"github.com/sparklee/abc-api/internal/data/ent/words"
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
	w, err := r.data.db.Words.
		Create().
		SetWord(word.Text).
		SetGroup("foo").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	word.Id = w.ID
	return word, nil
}

func (r *wordRepo) List(ctx context.Context, group string) ([]*biz.Word, error) {
	rows, err := r.data.db.Words.Query().Order(ent.Desc(words.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*biz.Word, 0)
	for _, word := range rows {
		result = append(result, &biz.Word{
			Id:    word.ID,
			Group: word.Group,
			Text:  word.Word,
		})
	}
	return result, nil
}

func (r *wordRepo) Delete(ctx context.Context, id int64) error {
	return r.data.db.Words.DeleteOneID(id).Exec(ctx)
}
