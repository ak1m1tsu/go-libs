// Package redis provides a redis connection.
package redis

import "github.com/redis/go-redis/v9"

// Connection represents a redis connection.
type Connection struct {
	*redis.Client
	cfg *Config
}

// Connect creates a new redis connection.
func Connect(url string, cfg *Config) (*Connection, error) {
	if cfg == nil {
		cfg = NewConfig()
	}

	conn := &Connection{cfg: cfg}

	clientOpts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	clientOpts.ReadTimeout = conn.cfg.readTimeout
	clientOpts.WriteTimeout = conn.cfg.writeTimeout
	clientOpts.DialTimeout = conn.cfg.dialTimeout

	conn.Client = redis.NewClient(clientOpts)

	return conn, nil
}
