package env

import (
	"github.com/kelseyhightower/envconfig"
)

type Values struct {
	DB
}

type DB struct {
	PostgresHost     string `default:"postgresql" split_words:"true"`
	PostgresPort     string `default:"5432" split_words:"true"`
	PostgresDatabase string `default:"connect" split_words:"true"`
	PostgresUser     string `required:"true" split_words:"true"`
	PostgresPassword string `required:"true" split_words:"true"`
}

func NewValue() (*Values, error) {
	var v Values
	err := envconfig.Process("", &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
