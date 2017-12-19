package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// RedisConfigHandler
type RedisConfigHandler struct {
	Conn redis.Conn
}

// Handle handle function from Handler interface
func (h RedisConfigHandler) Handle(name string) (val string, err error) {
	reply, err := h.Conn.Do("GET", name)
	if err != nil {
		return
	}

	val = fmt.Sprintf("%s", reply)
	return
}

// NewRedisConfigHandelr
func NewRedisConfigHandler(conn redis.Conn) *RedisConfigHandler {
	return &RedisConfigHandler{Conn: conn}

}
