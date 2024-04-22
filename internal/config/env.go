package config

import (
	"log"
	"os"
)

var envConfigMap map[string]string

func GetEnvVariable(envKey, defaultValue string) string {
	if envConfigMap == nil {
		envConfigMap = make(map[string]string)
	}

	if _, exist := envConfigMap[envKey]; exist {
		log.Printf("Returning cached value for key: %s", envKey)
		return envConfigMap[envKey]
	} else if v, exist := os.LookupEnv(envKey); exist && v != "" {
		envValue := os.Getenv(envKey)

		log.Printf("Saving new value in cache map for key: %s %s", envKey, envValue)
		envConfigMap[envKey] = envValue
		return envValue
	} else {
		log.Printf("Saving default value in cache map for key: %s", envKey)
		envConfigMap[envKey] = defaultValue

		return defaultValue
	}
}
