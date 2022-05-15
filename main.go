package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

var ctx = context.Background()

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	filesList := make([]string, 0, len(files))

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	key := 0
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			fullPath := filepath.Join(path, fileName)
			fmt.Println(fullPath)
			//fmt.Println(fileName, file.IsDir())
			filesList = append(filesList, fullPath)
			set(key, fullPath)
			fmt.Println(key)
			key++
		}
	}
}

func set(key int, value string) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis",
		DB:       0,
	})

	err := client.Set(
		ctx,
		strconv.Itoa(key),
		value,
		0).Err()

	if err != nil {
		log.Fatal(err)
	}
}
