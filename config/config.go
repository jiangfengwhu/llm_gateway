package config

type ServerConfig struct {
	T2IAddr string `json:"t2i_addr"`
}

const (
	T2I = "t2i_addr"
)

func UpdateT2IAddr(addr string) {
	Config.T2IAddr = addr
}
func GetServerConfig() ServerConfig {
	return Config
}

var Config ServerConfig
