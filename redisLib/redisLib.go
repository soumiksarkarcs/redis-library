package redisLib

import (
	"github.com/go-redis/redis"
)

/* To set a new record in redis
   Parameter 1: Represents key of the new entry
   Parameter 2: Represents the value of above key
   Parameter 3: Represents client connection to redis server
   Return type: nil if true;
   				else error
*/
func SetKey(key string, value interface{}, client *redis.Client) error {
	err := client.Set(key, value, 0).Err()
	return err
}

/* To get the value of an existing record in redis
   Parameter 1: Represents key of the record
   Parameter 2: Represents client connection to redis server
   Return type: (value, nil) if success;
                else error
*/
func GetKey(key string, client *redis.Client) (string, error) {
	val, err := client.Get(key).Result()
	return val, err
}

/* To get a new connection of redis server
   Parameter 1: Represents host address of redis server
   Return type: (client, nil) if success;
                else (nil, error)
*/
func GetConnection(hostAddress string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     hostAddress,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return client, nil
}
