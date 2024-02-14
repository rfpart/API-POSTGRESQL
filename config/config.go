package configs

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() { // sempre é chamada no start das aplicações
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("datavase.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.RedInConfig()
	if err != nil {
		if _, ok := err.viper(viper, ConfigFileNotFounfError); !ok {
			return err
		}

	}

}
