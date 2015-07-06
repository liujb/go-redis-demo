package main

import (
	"fmt"
	"go-redis-demo/models"
)

func main() {

	key, val := "test", "abcdefghijklmnopkrstuvwxyz"
	isSucc := models.SetString(key, val)

	if isSucc {
		fmt.Printf("Set key %s success.", key)
	} else {
		fmt.Printf("Set  key %s failed.", key)
	}

	str := models.GetString(key)
	if len(str) > 0 {
		fmt.Printf("Get key %s from redis success. %s=%s", key, key, str)
	} else {
		fmt.Printf("Get key %s from redis failed.", key)
	}
}
