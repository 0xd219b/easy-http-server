package config

const (
	RequestMethod = "request_method"
	ServerPort    = "server_port"
	RequestUrl    = "request_url"
)

type Config struct {
	RequestMethod string
	ServerPort    int64
	RequestUrl    string
}

var Cfg Config

func (c *Config) GetRequestMethod() string {
	return c.RequestMethod
}

func (c *Config) GetServerPort() int64 {
	return c.ServerPort
}

func (c *Config) GetRequestURL() string {
	return c.RequestUrl
}
