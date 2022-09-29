package service

import (
	"context"
	"exam/entity"
)

type Repository interface {
	InsertWord(ctx context.Context, words entity.Words) error
	UpdateWord(ctx context.Context, key string, add int32) error
	GetAllWord(ctx context.Context) ([]entity.Words, error)
	GetLimitedWord(ctx context.Context, page, limit int32) ([]entity.Words, error)
}
