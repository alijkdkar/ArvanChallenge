package discountrepository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	discountdomain "github.com/alijkdkar/ArvanChallenge/discount-domain"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type DiscountOpportunityRepository struct {
	RedisDb redis.Client
}

func CreateNewDiscountRepositoryInstance() *DiscountOpportunityRepository {
	return &DiscountOpportunityRepository{
		RedisDb: *RedisDb,
	}
}

// check Repository Can satisfy
var _ discountdomain.IDiscountOpportunityRepository = (*DiscountOpportunityRepository)(nil)

func (rp *DiscountOpportunityRepository) CreateDiscountOpp(disc discountdomain.DiscountOpportunity) error {
	ctx := context.Background()

	rp.RedisDb.HSet(ctx, "discount-Id:"+disc.Id.String(), disc.DiscountToMap())

	return nil
}

func (rp *DiscountOpportunityRepository) RemoveDiscountOpp(Id uuid.UUID) error {
	ctx := context.Background()
	res, er := rp.RedisDb.Exists(ctx, "discount-Id:"+Id.String()).Result()
	fmt.Println("delete Data Exists Check:", res, er)
	if res == 0 || er != nil {
		return errors.New("discount not fount")
	}
	_, err := rp.RedisDb.Del(ctx, "discount-Id:"+Id.String()).Result()
	if err != redis.Nil {
		return err
	}
	return nil
}

func (rp *DiscountOpportunityRepository) ExistsDiscountOpp(Id uuid.UUID) bool {
	ctx := context.Background()
	res, er := rp.RedisDb.Exists(ctx, "discount-Id:"+Id.String()).Result()
	fmt.Println("delete Data Exists Check:", res, er)
	if res == 0 || er != nil {
		return false
	}

	return true
}

func (rp *DiscountOpportunityRepository) AddNewDiscountUse(discTrans *discountdomain.DiscountTransaction) error {
	ctx := context.Background()
	discountKey := "discount-Id:" + discTrans.DiscountId.String()

	fmt.Println("in repositor trans:", discTrans)
	err := rp.RedisDb.Watch(ctx, func(tx *redis.Tx) error {
		fields := []string{"MaxCount", "UsedCount"}
		tran := rp.RedisDb.TxPipeline()
		counterRes, er := tx.HMGet(ctx, discountKey, fields...).Result()
		if er != nil {
			fmt.Println("GetCounter", er)
			return er
		}
		fmt.Println("after HMGet")
		fmt.Println(counterRes...)
		MaxCount, er := strconv.Atoi(counterRes[0].(string))
		UsedCount, er1 := strconv.Atoi(counterRes[1].(string))
		fmt.Println(er, er1, MaxCount, UsedCount)
		if er != nil || er1 != nil {
			return errors.New("internal server error 1000")
		}

		if MaxCount <= UsedCount {
			fmt.Println("discount ended")
			return errors.New("discount ended")
		}

		usedBefore, errr := tx.Exists(ctx, "d-tra-"+discTrans.DiscountId.String()+"-"+discTrans.CardId.String()).Result()

		if errr != nil {
			fmt.Println("check used opertion error ")
			return errors.New("server side error")
		}
		if usedBefore == 1 {
			fmt.Println(" its userd ")
			return errors.New("you have used of this code")
		}

		tx.HIncrBy(ctx, "discount-Id:"+discTrans.DiscountId.String(), "UsedCount", 1)
		tx.HSet(ctx, "d-tra-"+discTrans.DiscountId.String()+"-"+discTrans.CardId.String(), discTrans.DiscountTranssToMap())

		result, err := tran.Exec(ctx)
		if err != nil {
			return err
		}
		print(result)
		//todo :must throw integration Event to main Service
		return nil
	}, discountKey)
	if err != nil {
		return err
	}

	errPublish := rp.RedisDb.Publish(ctx, "chanel1", "Hello, World!").Err()
	if errPublish != nil {
		fmt.Println("Error publishing message:", errPublish)
	} else {
		fmt.Println("Message published successfully")
	}
	return nil
}
