package pubsub

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/pubsub", new(Redis))
}

// Redis is the k6 Redis extension
type Redis struct{}

func (Redis) Hello() string {
	return "Hello"
}
