package handler

import (
	"encoding/json"
	"log"
	"data-preparer/internal/service"

	"github.com/streadway/amqp"
)

type RabbitMQHandler struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	leagueSvc    *service.LeagueTableService
}

type Task struct {
	Task      string `json:"task"`
	LeagueID  string `json:"league_id,omitempty"`
	Timestamp string `json:"timestamp"`
}

func NewRabbitMQHandler(url string, leagueSvc *service.LeagueTableService) (*RabbitMQHandler, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"data_tasks",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQHandler{
		conn:      conn,
		channel:   ch,
		leagueSvc: leagueSvc,
	}, nil
}

func (h *RabbitMQHandler) Start() {
	msgs, err := h.channel.Consume(
		"data_tasks",
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var task Task
			if err := json.Unmarshal(d.Body, &task); err != nil {
				log.Printf("Error unmarshaling task: %v", err)
				continue
			}

			switch task.Task {
			case "update_league_tables":
				log.Println("Processing league tables update")
				h.leagueSvc.UpdateAllTables()
			default:
				log.Printf("Unknown task: %s", task.Task)
			}
		}
	}()

	log.Printf("RabbitMQ handler started")
	<-forever
}