package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault("cors.allow,methods", []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"})
	viper.SetDefault("cors.allow.headers", []string{"Content-Type", "Accept", "Authorization", "Origin", "Host"})
	viper.SetDefault("cors.allow.origins", "http://localhost")
}

func BindEnvVars() {
	bindEnv("cors.allow,methods", "CORS_ALLOW_METHODS")
	bindEnv("cors.allow.headers", "CORS_ALLOW_HEADERS")
	bindEnv("cors.allow.origins", "CORS_ALLOW_ORIGIN")
}

func PrepareConfig()  {
	SetDefaults()
	BindEnvVars()
}

func SetupViper(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".message-server-go" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".message-server-go"))
		viper.SetConfigName("auth-server")
		viper.SetConfigType("yml")
	}
	PrepareConfig()
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func bindEnv(key string, envVar string) {
	if err := viper.BindEnv(key, envVar); err != nil {
		log.Panic(fmt.Sprintf("Failed to bind config key '%s' to environment variable '%s'", key, envVar))
	}
}
