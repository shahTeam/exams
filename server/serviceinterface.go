package server

import (
	"context"
	"exam/entity"
)

type Service interface {
	InsertWords(ctx context.Context, words []entity.Words) error
	GetAllWords(ctx context.Context) ([]entity.Words, error)
	GetLimitedWords(ctx context.Context, page, limit int32) ([]entity.Words, error)
}
