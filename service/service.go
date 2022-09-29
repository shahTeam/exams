package service

import (
	"context"
	"exam/entity"
	//"fmt"
	"log"
)

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) InsertWords(ctx context.Context, words []entity.Words) error {
	wordsRepo, err := s.repo.GetAllWord(ctx)
	if err != nil {
		log.Println("error calling repository1 ", err)
	}
	if len(wordsRepo) == 0 {
		for i := 0; i < len(words); i++ {
			err = s.repo.InsertWord(context.Background(), words[i])
		}
		return err
	}
	count := 0
	for i := 0; i < len(words); i++ {
		count = 0
		for j := 0; j < len(wordsRepo); j++ {
			if words[i].Key == wordsRepo[j].Key {
				count += 1
			}
		}
		if count > 0 {
			err = s.repo.UpdateWord(ctx, words[i].Key, words[i].Value)
		} else {
			err = s.repo.InsertWord(context.Background(), words[i])
		}
	}
	return err
}
func (s Service) GetAllWords(ctx context.Context) ([]entity.Words, error) {
	words, err := s.repo.GetAllWord(ctx)
	if err != nil {
		log.Println("error while getting all words", err)
		return nil, err
	}
	return words, nil
}
func (s Service) GetLimitedWords(ctx context.Context, page, limit int32) ([]entity.Words, error) {
	newWordLists := make([]entity.Words, 0)
	wordlists, err := s.repo.GetLimitedWord(context.Background(), page, limit)
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	if int32(len(wordlists)) > offset + limit{
        for i := offset; i < offset+limit; i++ {
			newWordLists = append(newWordLists, wordlists[i])
		}
	}
	
	return newWordLists, nil
}
