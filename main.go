package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var (
	redisHost string
	retries   int
)

func init() {
	flag.StringVar(&redisHost, "redis", "127.0.0.1:6379", "Redis host:port")
	flag.IntVar(&retries, "retries", 100, "Maximum number of retries before giving up")
	flag.Parse()
}

func main() {
	var logger = log.New(os.Stdout, "", log.LstdFlags)
	redis.SetLogger(logger)

	client := redis.NewClient(&redis.Options{
		Addr:            redisHost,
		MaxRetries:      retries,
		MinRetryBackoff: 200 * time.Millisecond,
		MaxRetryBackoff: 5 * time.Second,
	})

	for {
		v, err := client.Info("clients").Result()
		if err != nil {
			logger.Fatal(err)
		}
		logger.Printf("\n%s---", v)
		time.Sleep(3 * time.Second)
	}
}
