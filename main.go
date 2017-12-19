package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/imantung/confiig"
)

var (
	Network = flag.String("network", "tcp", "redis network protocol")
	Address = flag.String("address", ":6379", "redis server address")
)

func main() {
	flag.Parse()

	// redis dial, you may use redis pool
	conn, err := redis.Dial(*Network, *Address)
	if err != nil {
		log.Fatal(err.Error())
	}

	// set redis handler
	confiig.SetHandler(NewRedisConfigHandler(conn))

	// register 'host' parameter
	confiig.Register("host")

	// get 'host' value
	host, err := confiig.Get("host")

	// some printing
	if err != nil {
		fmt.Printf("error - %s\n", err)
	} else {
		fmt.Printf("success - %s", host)
	}

}
