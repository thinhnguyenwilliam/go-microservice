package kafka

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

const (
	kafkaBroker = "localhost:9092"
	topic       = "test-topic"
	groupID     = "my-consumer-group"
)

func main() {
	// Start the consumer group listener in a goroutine
	go startConsumerGroup()

	// Set up Gin
	r := gin.Default()
	r.POST("/publish", publishHandler)
	r.Run(":8080")
}

func publishHandler(c *gin.Context) {
	var req struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	writer := kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("key"),
		Value: []byte(req.Message),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kafka write failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Message sent to Kafka"})
}

func startConsumerGroup() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{kafkaBroker},
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // Commit offsets periodically
	})
	defer r.Close()

	fmt.Println("Kafka consumer group started...")
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Read error: %v\n", err)
			continue
		}
		fmt.Printf("Received message: %s\n", string(msg.Value))
	}
}
