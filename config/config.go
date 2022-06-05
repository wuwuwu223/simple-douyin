package config

type Config struct {
	UseCos     bool   `json:"use_cos" mapstructure:"use_cos"`
	ListenPort int    `json:"listen_port" mapstructure:"listen_port"`
	BaseUrl    string `json:"base_url" mapstructure:"base_url"`
	Cos        struct {
		SecretID  string `json:"secret_id" mapstructure:"secret_id"`
		SecretKey string `json:"secret_key" mapstructure:"secret_key"`
		Address   string `json:"address" mapstructure:"address"`
	} `json:"cos" mapstructure:"cos"`
	Mysql struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"mysql" mapstructure:"mysql"`
}
