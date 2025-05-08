package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mousavisam/streaming-event-processor/internal/kafka"
	"github.com/mousavisam/streaming-event-processor/pkg/models"
)

type Handler struct {
	producer *kafka.Producer
}

func NewHandler(producer *kafka.Producer) *Handler {
	return &Handler{producer: producer}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST("/events", h.handlePostEvent)
}

func (h *Handler) handlePostEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event format"})
		return
	}

	// Set default fields if missing
	if event.CreatedAt.IsZero() {
		event.CreatedAt = time.Now()
	}

	// TODO: validate SchemaVersion, Data structure, etc.

	if err := h.producer.SendEvent(c.Request.Context(), &event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send event to Kafka"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event received and processed"})
}
