package configuration

import "github.com/spf13/viper"

type Configuration struct {
	MongoDBHost     string `mapstructure:"MONGODB_URI"`
	MongoDBDatabase string `mapstructure:"MONGODB_DATABASE"`
	WebServerPort   string `mapstructure:"PORT"`
	GraphQLPort     string `mapstructure:"GRAPHQL_PORT"`
	GrpcPort        string `mapstructure:"GRPC_PORT"`
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func (c *Configuration) Load(filename, ext, path string) (*Configuration, error) {

	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
