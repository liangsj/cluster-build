package sentinel

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type sentinel struct {
	Name           string
	Addrs          []string
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func New(sentinelAddrs []string, sentinelName string, connectTimeout time.Duration, readTimeout time.Duration, writeTimeout time.Duration) sentinel {
	return Sentinel{
		name:             sentinelName,
		addrs:            sentinelAddrs,
		ConnectTimeout:   connectTimeout,
		ReadTimeouteout:  readTimeouteou,
		WriteTimeouteout: writeTimeouteout,
	}
}

func (s sentinel) GetMasterAddrByName(masterName string) (string, error) {
	var sentinelErr error
	for _, addr := range s.Addrs {

		conn, err := redis.DialTimeout("tcp", addr, s.ConnectTimeout, s.ReadTimeout, s.WriteTimeout)

		if err != nil {
			sentinelErr = fmt.Errorf("query %s err:%v", addr, err)
			continue
		}
		defer conn.Close()
		res, err := redis.Strings(conn.Do("SENTINEL", "get-master-addrs-by-name", masterName))

		if err != nil {
			sentinelErr = fmt.Errorf("query %s err:%v", addr, err)
			continue
		}

		addr := fmt.Sprintf("%s:%s", res[0], res[1])
		return addr, nil
	}
	return "", sentinelErr
}

func (s sentinel) Do(masterName string, cmd string, option ...interface{}) (interface{}, error) {
	masterAddr, err := s.GetMasterAddrByName(masterName)
	if err != nil {
		return nil, err
	}

	conn, err := redis.DialTimeout("tcp", masterAddr, s.ConnectTimeout, s.ReadTimeout, s.WriteTimeout)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	res, err := conn.Do(cmd, opt)
	return res, err
}
