package api

// var client

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/util"
)

var client *redis.Client = nil

func initialConnectionToRedis() error {
	server_info, err := util.GetConfigs()
	if err != nil {
		return err
	}
	fmt.Println("Creating redis client connection to: ", server_info.REDIS_ADDR)
	client = redis.NewClient(&redis.Options{
		Addr:     server_info.REDIS_ADDR,
		Password: "",
		DB:       0,
	})
	return nil
}

// It returns shortened URL or error
func GetRedis(key string) (string, error) {
	if client == nil {
		err := initialConnectionToRedis()
		if err != nil {
			return "", err
		}
	}
	if client == nil {
		return "", errors.New("nil redis client")
	}
	if err := client.Ping(); err.Err() != nil {
		return "", errors.New("no redis client connection")
	}
	data := client.Get(key)
	if errors.Is(data.Err(), redis.Nil) {
		return "", redis.Nil
	}
	return data.Val(), nil
}

// It save shortened URL on redis
func SendRedis(key, val string) error {
	if client == nil {
		err := initialConnectionToRedis()
		if err != nil {
			return err
		}
	}
	if client == nil {
		return errors.New("nil redis client")
	}
	if err := client.Ping(); err.Err() != nil {
		return errors.New("no redis client connection")
	}
	var rt int = 0
	conf, err := util.GetConfigs()
	rt, err = strconv.Atoi(conf.REDIS_TIME)
	if err != nil {
		rt = 300
	}
	exp := time.Duration(time.Duration(rt) * time.Second) // 5 minutes is default
	cmd := client.Set(key, val, exp)
	return cmd.Err()
}
