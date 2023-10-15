package db

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AI1411/connectpj/internal/env"
)

type Client interface {
	Conn(ctx context.Context) *gorm.DB
}

type client struct {
	db *gorm.DB
}

func NewClient(e *env.DB) Client {
	db, _ := open(e)
	return &client{
		db: db,
	}
}

func open(env *env.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		env.PostgresHost,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDatabase,
		env.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *client) Conn(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}
