package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"

	"github.com/noskov-sergey/hh-auth/internal/domain"
	"github.com/noskov-sergey/hh-auth/internal/domain/agregates"
	"github.com/noskov-sergey/hh-auth/internal/infrastructure/hh"
	"github.com/noskov-sergey/hh-auth/internal/repository/auth/pgsql"
)

func main() {
	log, _ := zap.NewProduction()

	cl := hh.NewClient("https://api.hh.ru", "refresh_token")
	agr, err := cl.RefreshAccessToken(context.Background(), "USERNRQJ1ITF6VE632IBTN95DRO9D621SQCK0OTHNKDJO0OV2005NB2ETEGTPFP5", "USERI1PU50U9D79I19R8M321ESSR8N7SBJ9L3SSM1IPF206U4R98QHQ804RT2MGB")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(agr)

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

	r, err := agregates.NewAuthorization(
		agregates.AuthorizationParams{
			ID:           1,
			AccessToken:  "USERI1PU50U9D79I19R8M321ESSR8N7SBJ9L3SSM1IPF206U4R98QHQ804RT2MGB",
			RefreshToken: "USERNRQJ1ITF6VE632IBTN95DRO9D621SQCK0OTHNKDJO0OV2005NB2ETEGTPFP5",
			Status:       domain.Active,
			Expired:      time.Now().AddDate(0, 0, 14),
			Created:      time.Now(),
		},
	)
	if err != nil {
		log.Fatal("failed to create authorization object:", zap.Error(err))
	}

	dbCl := pgsql.NewAuthRepository(db)
	err = dbCl.Create(*r)
	if err != nil {
		log.Error("failed to save authorization object:", zap.Error(err))
	}

	g, err := dbCl.GetExpired(context.Background())
	if err != nil {
		if errors.Is(err, pgsql.ErrNoAuthorization) {
			fmt.Println("no authorization found")
		} else {
			log.Error("failed to load authorization object:", zap.Error(err))
		}
	}

	err = dbCl.MarkInactive(context.Background(), 2)
	if err != nil {
		log.Error("failed to mark object:", zap.Error(err))
	}

	fmt.Println(g)
}
