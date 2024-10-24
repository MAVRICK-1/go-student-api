package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" ENV:"ADDRESS" env-required:"true"`
}

//env-default:"production" is used to set the default value of the environment variable
// Config is capitalized so that it can be exported
type Config struct {

	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path"  env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {

		flags := flag.String("config", "", "path to the config file") // flag is used to parse the command line arguments

		fmt.Println(flags) // it will print the address of the flags
		fmt.Println(*flags) // it will print the value of the flags

		flag.Parse() // parse the command line arguments and set the value of flags
		//why parse ? because we want to get the value of the flag

		fmt.Println(flags) // it will print the address of the flags
		fmt.Println(*flags) // it will print the value of the flags

		configPath = *flags // get the value of the flag and set it to the configPath

		//but we alread set the value in the flags then we did flag.Parse()  it did not affect the value of flags and why did we use pointer ?

		//because flag.Parse() is used to parse the command line arguments and set the value of flags
		//if we did not use pointer then the value of flags will not be set

		//can you give a demonstration of this ?
		//yes, sure
		//if we did not use pointer then the value of flags will not be set
		//for example
		// flags := "config"
		// flag.Parse()
		// fmt.Println(flags) // it will print config
		// but if we use pointer
		// flags := flag.String("config", "", "path to the config file")
		// flag.Parse()
		// fmt.Println(*flags) // it will print the value of the flag
	}
	if configPath == "" {
		log.Fatalf("config file not found %s", configPath)
	}

	if _,err := os.Stat(configPath); os.IsNotExist(err) {  // check if the file exists or not _,err := os.Stat(configPath) it will return the file info and error here _ is used to ignore the file info ';' is used to separate the two statements
		log.Fatalf("config file not found %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg) // read the config file and set the value to the config variable //example of cleanenv.ReadConfig("config.yml", &config)
	//out put of the above example
	//configPath = "config.yml"
	//config = Config{Env: "production", StoragePath: "/var/lib/storage", HTTPServer: {Address: "localhost:8080"}}

	if err != nil {
		fmt.Println(cfg)
		log.Fatalf("failed to read config file %s", err)
	}
	
	return &cfg

	
	
}