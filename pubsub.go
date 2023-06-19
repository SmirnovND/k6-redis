package pubsub

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/pubsub", new(Redis))
}

type Message struct {
	ClientKey   string
	MessageText string
}

// Redis is the k6 Redis extension
type Redis struct {
	Client *redis.Client
}

type Config struct {
	Host string
	Port string
}

func (r *Redis) SetConfig(config map[string]interface{}) error {
	cf := &Config{}
	err := mapstructure.Decode(config, cf)
	if err != nil {
		return err
	}

	r.Client = redis.NewClient(&redis.Options{
		Addr: cf.Host + ":" + cf.Port,
	})

	return nil
}

func (r *Redis) Publish(message map[string]interface{}, chanel string) error {
	msg := &Message{}
	err := mapstructure.Decode(message, msg)
	if err != nil {
		return err
	}
	
	var ctx = context.Background()

	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := r.Client.Publish(ctx, chanel, payload).Err(); err != nil {
		return err
	}

	return nil
}
