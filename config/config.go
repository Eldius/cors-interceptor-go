package config

import (
	"sort"

	"github.com/spf13/viper"
)

func GetCORSAllowedMethods() []string {
	return viper.GetStringSlice("cors.allow.methods")
}

func GetCORSAllowedHeaders() []string {
	return viper.GetStringSlice("cors.allow.headers")
}

func GetCORSAllowedOrigin() []string {
	return viper.GetStringSlice("cors.allow.origins")
}

func IsOriginAllowed(host string) bool {
	origins := GetCORSAllowedOrigin()
	i := sort.SearchStrings(origins, host)
    return i < len(origins) && origins[i] == host
}
