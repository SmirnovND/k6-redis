package xk6_pubsub

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/redis", new(Redis))
}

type Message struct {
	ClientKey   string
	MessageText string
}

// Redis is the k6 Redis extension
type Redis struct{
	client *redis.Client
}

type Config struct {
	Host     string
	Port     string
}

func (r *Redis) SetConfig(cf *Config) {
	r.client = redis.NewClient(&redis.Options{
		Addr: cf.Host + ":" + cf.Port,
	})
}

func (r *Redis) Publish(Message Message, chanel string) error {
	var ctx = context.Background()

	payload, err := json.Marshal(Message)
	if err != nil {
		return err
	}

	if err := r.client.Publish(ctx, chanel, payload).Err(); err != nil {
		return err
	}
	
	return nil
}
