package configs

type Configs struct {
	MONGO_USER     string `mapstructure:"MONGO_USER"`
	MONGO_PASSWORD string `mapstructure:"MONGO_PASSWORD"`
	MONGO_HOST     string `mapstructure:"MONGO_HOST"`
	MONGO_PORT     string `mapstructure:"MONGO_PORT"`
	MONGO_DATABASE string `mapstructure:"MONGO_DATABASE"`
	HTTP_PORT      string `mapstructure:"HTTP_PORT"`
}

func GetConfig() *Configs {
	return &Configs{
		MONGO_USER:     "admin",        //os.Getenv("MONGO_USER"),
		MONGO_PASSWORD: "password",     //os.Getenv("MONGO_PASSWORD"),
		MONGO_HOST:     "localhost",    //os.Getenv("MONGO_HOST"),
		MONGO_PORT:     "27017",        //os.Getenv("MONGO_PORT"),
		MONGO_DATABASE: "skinaapis_db", //os.Getenv("MONGO_DATABASE"),
		HTTP_PORT:      "9098",         //os.Getenv("HTTP_PORT"),
	}
}
