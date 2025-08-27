package config

import (
	"log"
	"os"
	"strconv"
)

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == ""{
		log.Fatalf("%s enviroment variable is missing", key)
	}
	return os.Getenv(key)
}

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceURl() string{
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int{
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid",portStr)
	}
	return port
}
