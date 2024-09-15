package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oneweerachai/service1/internal/external"
	"go.uber.org/zap"

	"net/http"

	"github.com/oneweerachai/service1/pkg/models"

	"github.com/jmoiron/sqlx"
)

type API struct {
    DB            *sqlx.DB
    ExternalClient *external.Client
    Logger        *zap.Logger
}

func RegisterRoutes(router *gin.Engine, db *sqlx.DB, externalClient *external.Client, log *zap.Logger) {
    api := &API{
        DB:            db,
        ExternalClient: externalClient,
        Logger:        log,
    }

    v1 := router.Group("/api/v1")
    {
        v1.GET("/items", api.GetItems)
        v1.POST("/items", api.CreateItem)
    }
}

// GetItems handles GET /api/v1/items
func (api *API) GetItems(c *gin.Context) {
    var items []models.Item
    err := api.DB.Select(&items, "SELECT * FROM items")
    if err != nil {
        api.Logger.Error("Failed to fetch items", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
        return
    }

    c.JSON(http.StatusOK, items)
}

// CreateItem handles POST /api/v1/items
func (api *API) CreateItem(c *gin.Context) {
    var item models.Item
    if err := c.ShouldBindJSON(&item); err != nil {
        api.Logger.Warn("Invalid request body", zap.Error(err))
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Insert into database
    _, err := api.DB.NamedExec(`INSERT INTO items (name, description) VALUES (:name, :description)`, &item)
    if err != nil {
        api.Logger.Error("Failed to insert item", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
        return
    }

    // Optionally, interact with external API
    externalData, err := api.ExternalClient.GetSomeData("/data")
    if err != nil {
        api.Logger.Error("External API call failed", zap.Error(err))
        // Decide whether to fail the request or proceed
    } else {
        api.Logger.Info("Received external data", zap.Any("data", externalData))
    }

    c.JSON(http.StatusCreated, item)
}
