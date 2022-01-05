package config

const (
	RequestMethod = "request_method"
	ServerPort    = "server_port"
	RequestUrl    = "request_url"
	MockfilePath  = "mockfile_path"
	MockService   = "mock_service"
)

type Config struct {
	RequestMethod string
	ServerPort    int64
	RequestUrl    string
	Mock          Mock
}

type Mock struct {
	FilePath    string
	ServiceType string
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

func (c *Config) GetMockfilePath() string {
	return c.Mock.FilePath
}

func (c *Config) GetMockService() string {
	return c.Mock.ServiceType
}
