package redis

import "time"

const (
	defaultReadTimeout  = time.Second * 5
	defaultWriteTimeout = time.Second * 5
	defaultDialTimeout  = time.Second * 5
)

// Config represents a configuration for redis client.
type Config struct {
	readTimeout  time.Duration
	writeTimeout time.Duration
	dialTimeout  time.Duration
}

// NewConfig creates a config for redis client.
//
//	Defaults:
//	  - read timeout = 5s
//	  - write timeout = 5s
//	  - idle timeout = 30s
func NewConfig() *Config {
	return &Config{
		readTimeout:  defaultReadTimeout,
		writeTimeout: defaultWriteTimeout,
		dialTimeout:  defaultDialTimeout,
	}
}

// WithReadTimeout sets the read timeout for redis client.
func (c *Config) WithReadTimeout(t time.Duration) *Config {
	c.readTimeout = t
	return c
}

// WithWriteTimeout sets the write timeout for redis client.
func (c *Config) WithWriteTimeout(t time.Duration) *Config {
	c.writeTimeout = t
	return c
}

// WithDialTimeout sets the dial timeout for redis client.
func (c *Config) WithDialTimeout(t time.Duration) *Config {
	c.dialTimeout = t
	return c
}
