package configs

import "github.com/spf13/viper"


type conf struct {
	PortServer string `mapstructure:"PORT_SERVER"`
	DBUser string `mapstructure:"DB_USER"`
	DBDatabase string  `mapstructure:"DB_DATABASE"`
	DBPassword string  `mapstructure:"DB_PASSWORD"`
	DBPort string  `mapstructure:"DB_PORT"`
	DBHost string  `mapstructure:"DB_HOST"`
	JwtSecret string  `mapstructure:"JWT_SECRET"`
	JwtExpiresIn string  `mapstructure:"JWT_EXPIRES_IN"`
}


func LoadEnv() (*conf,error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath("path")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if(err != nil){panic(err)}

	err = viper.Unmarshal(&cfg)
	if(err != nil) {panic(err)}

	return cfg, err


}