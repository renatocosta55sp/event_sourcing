package depositfunds

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/renatocosta55sp/event_sourcing/internal/domain/commands"
	"github.com/renatocosta55sp/event_sourcing/internal/infra/adapters/persistence"
	"github.com/renatocosta55sp/modeling/infra/bus"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	Db                  *pgxpool.Pool
	DomainEventRegistry *bus.EventRegistry
}

const requestDataKey = "requestData"

type BankAccountRequest struct {
	Id     string
	Amount float64 `json:"amount"  binding:"required"`
}

func ValidateRequest(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorD": "ID parameter is required"})
		ctx.Abort()
		return
	}

	var requestData BankAccountRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	requestData.Id = id

	ctx.Set("requestData", requestData)

	ctx.Next()

}

func (h HttpServer) DepositFunds(ctx *gin.Context) {

	requestData, exists := ctx.Get(requestDataKey)
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Request data not found"})
		return
	}

	data := requestData.(BankAccountRequest)
	aggregateIdentifier, err := uuid.Parse(data.Id)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	conn, err := h.Db.Acquire(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	commandResult, _, err := CommandExecutor{
		persistence.NewPersistentEventStore(conn, h.DomainEventRegistry, "public"),
		persistence.NewPersistentSnapshotEventStore(conn, h.DomainEventRegistry, "public"),
		persistence.TransactionDb{
			Ctx:  ctx,
			Conn: conn,
		},
	}.Send(
		ctx,
		commands.DepositFundsCommand{
			AggregateID: aggregateIdentifier,
			Amount:      data.Amount,
		},
	)

	if err != nil {
		logrus.WithError(err).Error("failed to validate bank account on deposit funds command")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"result": commandResult})

}
