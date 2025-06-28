package main

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"

	a "github.com/noskov-sergey/hh-auth/internal/controller/auth"
	"github.com/noskov-sergey/hh-auth/internal/infrastructure/hh"
	"github.com/noskov-sergey/hh-auth/internal/repository/auth/pgsql"
	"github.com/noskov-sergey/hh-auth/internal/service/auth"
)

func main() {
	log, _ := zap.NewProduction()

	db, err := sql.Open("postgres", "host=localhost port=5438 dbname=auth user=auth-user password=auth-password sslmode=disable")
	if err != nil {
		log.Error("failed to connect to database:", zap.Error(err))
		panic(err)
	}
	defer db.Close()

	err = goose.SetDialect("postgres")
	if err != nil {
		log.Error("failed to set postgres dialect:", zap.Error(err))
	}

	err = goose.Up(db, "migrations/")
	if err != nil {
		log.Error("failed to up migrations:", zap.Error(err))
	}

	dbCl := pgsql.NewAuthRepository(db)
	cl := hh.NewClient("https://api.hh.ru", "refresh_token")
	svc := auth.NewAuthService(dbCl, cl)

	a.NewRefresh(svc, log).Refresh(context.Background())
}
