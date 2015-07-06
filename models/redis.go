package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go-redis-demo/helpers"
)

func SetString(key, val string) bool {
	conn := helpers.GetConnect()

	if len(key) <= 0 || len(val) <= 0 {
		fmt.Printf("key and val must be not empty.")
		return false
	}

	_, err := conn.Do("SET", key, val)
	if err != nil {
		fmt.Printf("set %s=%s failed.", key, val)
		return false
	}

	defer conn.Close()
	return true
}

func GetString(key string) string {
	conn := helpers.GetConnect()

	if len(key) <= 0 {
		fmt.Printf("key must be not empty.")
		return ""
	}

	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Printf("get %s failed.", key)
		return ""
	}

	defer conn.Close()
	return val
}
