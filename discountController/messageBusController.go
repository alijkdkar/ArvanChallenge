package discountcontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	discountrepository "github.com/alijkdkar/ArvanChallenge/discount-repository"
	"github.com/google/uuid"
)

// register this service endpoint
func EnableMessageBusServices() {
	ProducerTransaction()
	go ConsumeComplite()
}

// this method produce unComplited transaction to core service every 3 second
func ProducerTransaction() {

	_dicRepo := discountrepository.CreateNewDiscountRepositoryInstance()
	ticker := time.NewTicker(3 * time.Second)

	checker := func() {
		data, er := _dicRepo.GetUnComplitedTransaction()
		if er != nil {
			fmt.Println("Get Uncomplite Trans error:", er)
		}

		_dicRepo.PublishToMessageBus(data)

	}

	go func() {
		for {
			select {
			case <-ticker.C:
				checker()
			}
		}
	}()
}

func ConsumeComplite() {
	_dicRepo := discountrepository.CreateNewDiscountRepositoryInstance()
	pubsub := _dicRepo.RedisDb.Subscribe(context.Background(), "Transaction-complite-channel")
	defer pubsub.Close()

	// Wait for messages
	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			fmt.Println("Error receiving message:", err)
			return
		}
		fmt.Printf("Received message on complite trans: %s\n", msg.Payload)
		var coomand TransCompliteDto
		er := json.Unmarshal([]byte(msg.Payload), &coomand)
		if er != nil {
			fmt.Println("error on un marshal complite command")
			continue
		}
		errcom := _dicRepo.CompliteTransaction(coomand.Id)
		if errcom != nil {
			fmt.Println("error on complite Trans")
			continue
		}

	}

}

type TransCompliteDto struct {
	Id uuid.UUID `json:"Id"`
}
