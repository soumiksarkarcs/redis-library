package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"spicasys.com/redis/redisLib"
)

// var (
// 	ErrNil    = errors.New("no matching record found in redis database")
// )

var (
	client    *redis.Client
	err       error
	RedisAddr = "localhost:6379"
)

func main() {
	fmt.Println("Go Redis Library :: ")

	for {
		fmt.Println("\n1] Create Connection\n2] Add Record\n3] Get Record\n4] Exit\n")
		var ch int
		fmt.Printf("Enter your choice : ")
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			client, err = redisLib.GetConnection(RedisAddr)
			if err != nil {
				fmt.Println("Failed to connect to redis: %s", err.Error())
				os.Exit(1)
			} else {
				fmt.Println("Connection created successfully")
			}
			//break
		case 2:
			var key, value string
			fmt.Printf("Enter key : ")
			fmt.Scanln(&key)
			fmt.Printf("Enter value : ")
			fmt.Scanln(&value)

			err = redisLib.SetKey(key, value, client)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Record added successfully")
			}
			//break
		case 3:
			var key_search string
			fmt.Printf("Enter key for search : ")
			fmt.Scanln(&key_search)

			value, err_get := redisLib.GetKey(key_search, client)
			if err_get != nil {
				fmt.Println(err_get)
				os.Exit(2)
			} else {
				fmt.Printf("%s --> %s\n", key_search, value)
			}
		case 4:
			os.Exit(0)
		}
	}
}
