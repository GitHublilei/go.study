package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func initRedis() error {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "182.61.13.234:6379",
		Password: "1198072529",
		DB:       0,
	})
	_, err := redisDB.Ping().Result()
	return err
}

func redisExample() {
	err := redisDB.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	val, err := redisDB.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := redisDB.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name dose not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

func redisZset() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "Javascript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// 把元素都追加到key
	num, err := redisDB.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println(num)

	newScore, err := redisDB.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分速最高的3个
	ret, err := redisDB.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95～100
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisDB.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connet redis failed, err:%v\n", err)
		return
	}
	// redisExample()
	redisZset()
}
