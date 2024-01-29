package postgresql

import "time"

const (
	defaultConnectionTimeout  = time.Second * 5
	defaultRetryDelay         = time.Second
	defaultConnectionAttempts = 5
	defaultPoolSize           = 10
)

// Config represents configuration for postgresql connection.
type Config struct {
	connectionTimeout  time.Duration
	retryDelay         time.Duration
	connectionAttempts uint
	poolSize           int32
}

// NewConfig creates a config for postgresql connection.
//
//	Defaults:
//	 - connection attempts = 5
//	 - connection timeout = 5s
//	 - pool size = 10
func NewConfig() *Config {
	return &Config{
		connectionAttempts: defaultConnectionAttempts,
		connectionTimeout:  defaultConnectionTimeout,
		poolSize:           defaultPoolSize,
		retryDelay:         defaultRetryDelay,
	}
}

// WithConnectionTimeout returns config with provided connection timeout.
func (c *Config) WithConnectionTimeout(timeout time.Duration) *Config {
	c.connectionTimeout = timeout
	return c
}

// WithConnectionAttempts returns config with provided connection attempts.
func (c *Config) WithConnectionAttempts(attempts uint) *Config {
	c.connectionAttempts = attempts
	return c
}

// WithPoolSize returns config with provided pool size.
func (c *Config) WithPoolSize(size int32) *Config {
	c.poolSize = size
	return c
}

// WithRetryDelay returns config with provided retry delay.
func (c *Config) WithRetryDelay(delay time.Duration) *Config {
	c.retryDelay = delay
	return c
}
