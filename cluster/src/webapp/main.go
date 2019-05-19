package main

import (
	"fmt"
	"lesson3/src/webapp/lns"
	"log"
	"net/http"
	"time"
	//使用开源的redigo 客户端
	"github.com/gomodule/redigo/redis"
)

type Sentinel struct {
	name  string
	Addrs []string
}

var S Sentinel = Sentinel{
	name:  "default",
	Addrs: []string{"redis_sentinel_1:26379", "redis_sentinel_2:26379"},
}

func main() {
	lns.RegisterServer()
	http.HandleFunc("/get", getParams)
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/setredis", set2Redis)
	http.HandleFunc("/getredis", getFromRedis)
	log.Print("start to liste 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Print(fmt.Sprintf("%v\n", err))
	}
	log.Print("server end")
}

//启动http服务器的测试
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HelloWold\n"))
}

func getParams(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	log.Println(fmt.Sprintf("%v\n", vars))

}

func set2Redis(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	if _, ok := vars["key"]; !ok {
		log.Println("querystring key is not exist")
	}

	if _, ok := vars["val"]; !ok {
		log.Println("querystring val is not exist")
	}

	conn, err := S.getRedisConn()

	if err != nil {
		log.Println(fmt.Sprintf("%v\n", err))
	}

	defer conn.Close()

	_, err = redis.String(conn.Do("set", vars["key"], vars["val"]))

	if err != nil {
		log.Println(fmt.Sprintf("%v\n", err))

	}

}

func getFromRedis(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	if _, ok := vars["key"]; !ok {
		log.Println("querystring key is not exist")
	}

	conn, err := S.getRedisConn()
	if err != nil {
		log.Println(fmt.Sprintf("%v\n", err))
	}

	defer conn.Close()
	reqStr, err := redis.String(conn.Do("get", vars["key"]))

	if err != nil {
		log.Println(fmt.Sprintf("%v\n", err))
	}
	w.Write([]byte(reqStr))
}

func (s *Sentinel) getRedisConn() (redis.Conn, error) {

	for _, addr := range s.Addrs {
		sentinelConn, err := redis.DialTimeout("tcp", addr, 0, 1*time.Second, 1*time.Second)
		if err != nil {
			continue
		}
		defer sentinelConn.Close()
		res, err := redis.Strings(sentinelConn.Do("SENTINEL", "get-master-addr-by-name", "master"))
		fmt.Printf("redis_master_addr:%v", res)
		if err != nil {
			return nil, err
		}

		redisConn, err := redis.DialTimeout("tcp", fmt.Sprintf("%s:%s", res[0], res[1]), 0, 1*time.Second, 1*time.Second)
		if err != nil {
			continue
		}
		return redisConn, nil

	}
	return nil, fmt.Errorf("sentinel err")
}
