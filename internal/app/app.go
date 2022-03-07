package app

import (
	"context"
	"github.com/core-go/health"
	q "github.com/core-go/sql"
	_ "github.com/lib/pq"

	. "search-users/internal/usecase/users"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	UserHandler   UserHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	db, err := q.OpenByConfig(root.Sql)
	if err != nil {
		return nil, err
	}

	userService := NewUserService(db)
	userHandler := NewUserHandler(userService)

	sqlChecker := q.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}