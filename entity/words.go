package entity

import "github.com/google/uuid"

type Words struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value int32  `json:"value"`
}
type WordList []Words

func NewWord(key string, value int32) Words {
	return Words{
		Id:    uuid.NewString(),
		Key:   key,
		Value: value,
	}
}
