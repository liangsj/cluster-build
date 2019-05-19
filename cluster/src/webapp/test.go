package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:7005")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	res, err := redis.Strings(conn.Do("SENTINEL", "get-master-addr-by-name", "master"))

	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(res)
}
