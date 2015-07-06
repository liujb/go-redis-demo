package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type RedisNode struct {
	Host string
	Port int
}

type RedisNodes struct {
	Examples []*RedisNode
}

/**
 * 获得redis节点
 */
func getNode() *RedisNode {

	var res *RedisNode
	bytes, err := ioutil.ReadFile("conf/redis.json")
	if err != nil {
		fmt.Printf("Read redis.json failed, error: ", err)
		return res
	}

	var nodes RedisNodes
	errors := json.Unmarshal(bytes, &nodes)
	if errors != nil {
		fmt.Printf("json.Unmarshal failed, error: ", errors)
		return res
	}

	length := len(nodes.Examples)
	if length <= 0 {
		fmt.Printf("Error.")
		return res
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(length)
	res = nodes.Examples[index]
	return res
}

/**
 * 获得redis.Conn
 * 操作完之后要 conn.Close()
 */
func GetConnect() redis.Conn {
	node := getNode()

	if node == nil {
		var res redis.Conn
		return res
	}

	conn, err := redis.Dial("tcp", node.Host+":"+strconv.Itoa(node.Port))
	if err != nil {
		fmt.Printf("Connect %s:%d error.", node.Host, node.Port, err)
	}

	return conn
}
