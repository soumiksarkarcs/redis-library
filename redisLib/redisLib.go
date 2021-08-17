package redisLib

import (
	"github.com/go-redis/redis"
)

type Redis interface {
	HSetKey(string, string, string) error
	HGetKey(string, string) *redis.StringCmd
	SetKey(string, interface{}) error
	GetKey(string) (string, error)
	GetConnection(string) error
}

type RedisClient struct {
	client *redis.Client
}

/* To set a new record in redis
   Parameter 1: Represents  name of the hash
   Parameter 2: Represents one of the field of hash
   Parameter 3: Represents value of the specified field of hash
   Parameter 4: Represents client connection to redis server
   Return type: nil if true;
   				else error
*/
func (c *RedisClient) HSetKey(hash string, field string, value string) error {
	err := c.client.HSet(hash, field, value).Err()
	return err
}

/* To get the field value of an existing hash in redis
   Parameter 1: Represents  name of the hash
   Parameter 2: Represents one of the field of hash
   Parameter 3: Represents client connection to redis server
   Return type: *redis.StringCmd
*/
func (c *RedisClient) HGetKey(hash string, field string) *redis.StringCmd {
	cmd := c.client.HGet(hash, field)
	return cmd
}

/* To set a new record in redis
   Parameter 1: Represents key of the new entry
   Parameter 2: Represents the value of above key
   Parameter 3: Represents client connection to redis server
   Return type: nil if true;
   				else error
*/
func (c *RedisClient) SetKey(key string, value interface{}) error {
	err := c.client.Set(key, value, 0).Err()
	return err
}

/* To get the value of an existing record in redis
   Parameter 1: Represents key of the record
   Parameter 2: Represents client connection to redis server
   Return type: (value, nil) if success;
                else error
*/
func (c *RedisClient) GetKey(key string) (string, error) {
	val, err := c.client.Get(key).Result()
	return val, err
}

/* To get a new connection of redis server
   Parameter 1: Represents host address of redis server
   Return type: (client, nil) if success;
                else (nil, error)
*/
func (c *RedisClient) GetConnection(hostAddress string) error {
	c.client = redis.NewClient(&redis.Options{
		Addr:     hostAddress,
		Password: "",
		DB:       0,
	})
	if err := c.client.Ping().Err(); err != nil {
		return err
	}
	return nil
}
