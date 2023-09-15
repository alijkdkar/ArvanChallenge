package discountrepository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

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

	return rp.RedisDb.Watch(ctx, func(tx *redis.Tx) error {
		fields := []string{"MaxCount", "UsedCount", "Amount"}
		tranId := "d-tra-" + discTrans.Id.String()
		tran := rp.RedisDb.TxPipeline()

		counterRes, er := tx.HMGet(ctx, discountKey, fields...).Result()
		if er != nil {
			fmt.Println("GetCounter", er)
			return er
		}

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

		usedBefore, errr := tx.Exists(ctx, tranId).Result()

		if errr != nil {
			fmt.Println("check used opertion error ")
			return errors.New("server side error")
		}
		if usedBefore == 1 {
			fmt.Println(" its userd ")
			return errors.New("you have used of this code")
		}
		AmountOfDiscount, er1 := strconv.Atoi(counterRes[2].(string))
		if er1 != nil {
			return errors.New("discount amount not valid")
		}
		discTrans.Amount = int64(AmountOfDiscount)
		tx.HIncrBy(ctx, discountKey, "UsedCount", 1)
		discTrans.ChangeStatus(discountdomain.Saved)
		tx.HSet(ctx, tranId, discTrans.DiscountTranssToMap())

		result, err := tran.Exec(ctx)
		if err != nil {
			fmt.Println("error on transaction exec:", err)
			return err
		}
		fmt.Println("exec trans resualt:", result)

		return nil
	}, discountKey)

}

func (rp *DiscountOpportunityRepository) GetUnComplitedTransaction() ([]*discountdomain.DiscountTransaction, error) {

	fmt.Println("in Gets")
	ctx := context.Background()
	res := []*discountdomain.DiscountTransaction{}
	keys, err := rp.RedisDb.Keys(ctx, "*d-tra-*").Result()

	if err != nil {
		return nil, err
	}

	for _, v := range keys {

		tran, er := rp.RedisDb.HGetAll(ctx, v).Result()
		if er != nil {
			fmt.Println("key Read Error", v, er)
			continue
		}
		status, errr := strconv.Atoi(tran["Status"])
		if errr != nil {
			return nil, errr
		}
		if discountdomain.Saved == discountdomain.TransactionStatus(status) {
			tranTemp := discountdomain.DiscountTransaction{}
			tranTemp.LoadFromMap(tran)
			res = append(res, &tranTemp)
		}
	}
	return res, nil
}

func (rp *DiscountOpportunityRepository) CompliteTransaction(transId uuid.UUID) error {

	ctx := context.Background()

	res, er := rp.RedisDb.HSet(ctx, "d-tra-"+transId.String(), "Status", fmt.Sprintf("%d", discountdomain.Complited)).Result()
	if er != nil {
		fmt.Println("error on Save complite in db")
	}
	fmt.Println("res of set status of complit", res)

	return nil
}

func (rp *DiscountOpportunityRepository) PublishToMessageBus(data []*discountdomain.DiscountTransaction) {
	// this method recive data and publish to message bus and save in cash until one number
	ctx := context.Background()
	fmt.Println("Publish data ......", len(data))
	for _, v := range data {
		if d, er := rp.RedisDb.Exists(ctx, "p-b-m-"+v.Id.String()).Result(); er == nil {

			if d == 0 {
				//set in db with expire time about 30 seceand

				fmt.Println("message Published")
				js, e := json.Marshal(v)
				if e != nil {
					fmt.Println("error on marshal", e)
				}
				res, errrr := rp.RedisDb.Publish(ctx, "disc-tran-publish", js).Result()
				if errrr != nil {
					fmt.Println("error on publish ", errrr)
				}
				rp.RedisDb.Set(ctx, "p-b-m-"+v.Id.String(), "-", time.Second*30)
				fmt.Println("mesage publish result", res)
			} else {
				fmt.Println("its have had sent ")
			}
		}

	}

}
