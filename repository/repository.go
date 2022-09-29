package repository

import (
	"context"
	//"exam/config"
	//"exam/connection"
	"exam/entity"
	"fmt"
	"log"
      "database/sql"
	//"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p Postgres) InsertWord(ctx context.Context, words entity.Words) error {
	query := `INSERT INTO words VALUES ($1,$2,$3)`
	if _, err := p.db.ExecContext(ctx, query, words.Id, words.Key, words.Value); err != nil {
		log.Println("error while inserting repo", err)
		return err
	}
	return nil
}
func (p Postgres) UpdateWord(ctx context.Context, key string, add int32) error {
	query := `UPDATE words SET value = value + $1 WHERE keys = $2`
	if _, err := p.db.ExecContext(ctx, query, add, key); err != nil {
		return err
	}
	return nil
}
func (p Postgres) GetAllWord(ctx context.Context) ([]entity.Words, error) {
	words := make([]entity.Words, 0)
	query := `SELECT * FROM words ORDER by vVALUE DESC`
	fmt.Println(p.db)
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("error while reading from repository", err)
		return nil, err
	}
	for rows.Next() {
		var word entity.Words
		if err := rows.Scan(&word.Id, &word.Key, &word.Value); err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}
func (p Postgres) GetLimitedWord(ctx context.Context, page, limit int32) ([]entity.Words, error) {
	words, err := p.GetAllWord(ctx)
	if err != nil {
		return []entity.Words{}, err
	}
	return words, nil
}
