package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	discountrepository "github.com/alijkdkar/ArvanChallenge/discount-repository"
	"github.com/alijkdkar/ArvanChallenge/domain"
	"github.com/alijkdkar/ArvanChallenge/repository"
	"github.com/google/uuid"
)

func EnableCoreMessageBusServices() {
	ConsumeIntegrtedTran()

}

func ConsumeIntegrtedTran() {
	_dicRepo := discountrepository.CreateNewDiscountRepositoryInstance()
	_carRep := repository.NewCreditCardRepository()
	fmt.Println("in Ricever")
	go func() {
		// Subscribe to a channel
		pubsub := _dicRepo.RedisDb.Subscribe(context.Background(), "disc-tran-publish")
		defer pubsub.Close()

		// Wait for messages
		for {
			msg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				fmt.Println("Error receiving message:", err)
				return
			}
			// fmt.Printf(" ************Received message: %s\n", msg.Payload)
			var newdisc DiscountTransactionDto
			err1 := json.Unmarshal([]byte(msg.Payload), &newdisc)
			if err1 != nil {
				fmt.Println("eror on cast message recive", err)
				continue
			}
			tr := domain.CreateNewTransaction(float64(newdisc.Amount), 1, newdisc.CardId, true, newdisc.DiscountId)

			saveEr := _carRep.AddTransaction(tr)
			if saveEr != nil {
				fmt.Println("error on save transaction from message bus:", saveEr)
			}
			//	must publish to
			fmt.Println("publish Complite Trans from core to disc")
			senInstance := CompliteSendTransDto{Id: newdisc.Id}
			js, erCast := json.Marshal(senInstance)
			if erCast != nil {
				fmt.Println("error on murshal sent instance")
				continue
			}
			sent, publishError := _dicRepo.RedisDb.Publish(context.Background(), "Transaction-complite-channel", js).Result()
			if publishError != nil {
				fmt.Println("error on publish complite flag to core :", publishError)

			}
			fmt.Println("result af respone flag :", sent)
		}
	}()
}

type DiscountTransactionDto struct {
	Id         uuid.UUID
	DiscountId uuid.UUID
	CardId     uuid.UUID
	CreatedAt  time.Time
	Status     int
	Amount     int64
}

type CompliteSendTransDto struct {
	Id uuid.UUID `json:"Id"`
}
