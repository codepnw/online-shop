package auth

import (
	"context"
	"testing"

	"github.com/codepnw/online-shop/external/database"
	"github.com/codepnw/online-shop/internal/config"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    "codepnw@mail.com",
		Password: "mypassword",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}
