package infra

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/renatocosta55sp/event_sourcing/internal/slices/depositfunds"
	"github.com/renatocosta55sp/modeling/infra/bus"
)

func InitRoutes(
	r *gin.RouterGroup,
	db *pgxpool.Pool,
	domainEventRegistry *bus.EventRegistry) {

	resDepositFunds := depositfunds.HttpServer{Db: db, DomainEventRegistry: domainEventRegistry}
	r.PATCH("/accounts/:id/deposit", depositfunds.ValidateRequest, resDepositFunds.DepositFunds)
}
