package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

//Driver struct
type Driver struct {
	DSN                 []string
	InternalPoolTimeout int64
	IdleTimeout         int64
	ReadTimeout         int64
	WriteTimeout        int64
	PingTry             bool

	conn redis.UniversalClient
}

func NewRedis(dsn []string, internalPoolTimeout, idleTimeout, readTimeout, writeTimeout int64) (*Driver, error) {
	driver := &Driver{
		DSN:                 dsn,
		InternalPoolTimeout: internalPoolTimeout,
		IdleTimeout:         idleTimeout,
		ReadTimeout:         readTimeout,
		WriteTimeout:        writeTimeout,
	}

	err := driver.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return driver, nil
}

//Connect new redis
func (r *Driver) Connect() error {
	log.Println("Starting the connection to the redis...")
	// create redis universal client
	r.conn = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        r.DSN,
		PoolTimeout:  time.Duration(r.InternalPoolTimeout) * time.Second,
		IdleTimeout:  time.Duration(r.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(r.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(r.WriteTimeout) * time.Second,
	})

	// first ping
	_, err := r.conn.Ping().Result()
	if err != nil {
		return err
	}

	// health ping
	if r.PingTry {
		r.healthPing()
	}

	log.Println("Connected to the Redis")

	return nil
}

//Conn get connection
func (r *Driver) Conn() redis.UniversalClient {
	return r.conn
}

//healthPing start health ping
func (r *Driver) healthPing() {
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				_, err := r.conn.Ping().Result()
				if err != nil {
					<-quit
					r.Connect()
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
