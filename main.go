package main

import (
	"fmt"
	"os"

	"spicasys.com/redis/redisLib"
)

var (
	client    redisLib.RedisClient
	err       error
	RedisAddr = "localhost:6379"
)

func main() {
	fmt.Println("Go Redis Library :: ")
	for {
		fmt.Printf("\n1] Create Connection\n2] SET\n3] GET\n4] HSET\n5] HGET\n6] Exit\n")
		var ch int
		fmt.Printf("Enter your choice : ")
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			err = client.GetConnection(RedisAddr)
			if err != nil {
				fmt.Printf("Failed to connect to redis: %s", err.Error())
				os.Exit(1)
			} else {
				fmt.Println("Connection created successfully")
			}
		case 2:
			var key, value string
			fmt.Printf("Enter key : ")
			fmt.Scanln(&key)
			fmt.Printf("Enter value : ")
			fmt.Scanln(&value)

			err = client.SetKey(key, value)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Record added successfully")
			}
		case 3:
			var key_search string
			fmt.Printf("Enter key for search : ")
			fmt.Scanln(&key_search)

			value, err_get := client.GetKey(key_search)
			if err_get != nil {
				fmt.Println(err_get)
				os.Exit(2)
			} else {
				fmt.Printf("%s --> %s\n", key_search, value)
			}
		case 4:
			var hash, field, value string
			fmt.Printf("Enter hash name : ")
			fmt.Scanln(&hash)
			fmt.Printf("Enter field of hash : ")
			fmt.Scanln(&field)
			fmt.Printf("Enter value : ")
			fmt.Scanln(&value)

			err = client.HSetKey(hash, field, value)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Record added successfully")
			}
		case 5:
			var hash, field string
			fmt.Printf("Enter hash name : ")
			fmt.Scanln(&hash)
			fmt.Printf("Enter field of hash to search : ")
			fmt.Scanln(&field)

			cmd := client.HGetKey(hash, field)
			if cmd.Err() != nil {
				fmt.Println(cmd.Err())
				os.Exit(2)
			} else {
				fmt.Printf("%s --> %s --> %s\n", hash, field, cmd.Val())
			}
		case 6:
			os.Exit(0)
		}
	}
}
