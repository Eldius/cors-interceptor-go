package cors

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

const (
	originLocalhost  = "http://localhost"
	originNotAllowed = "http://localhost:8000"
)

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func init() {
	//config.PrepareConfig()
	viper.SetDefault("cors.allow.methods", []string{"POST", "GET"})
	viper.SetDefault("cors.allow.headers", []string{"Content-Type", "Accept", "Authorization", "Origin", "Host"})
	viper.SetDefault("cors.allow.origins", []string{originLocalhost})
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func TestRequest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler())
	s := httptest.NewServer(CORS(mux))
	defer s.Close()

	req, err := http.NewRequest(http.MethodGet, s.URL, &bytes.Reader{})
	if err != nil {
		t.Errorf("---\nFailed to create test request\n%s\n---", err.Error())
	}

	req.Header.Add("Origin", originLocalhost)
	c := http.Client{}

	res, err := c.Do(req)
	if err != nil {
		t.Errorf("---\nFailed to execute test request\n%s\n---", err.Error())
	}

	if res.Header.Get("Access-Control-Allow-Origin") == "" {
		t.Errorf("Response should have 'Access-Control-Allow-Origin' header, but it hasn't")
	}

	if res.Header.Get("Access-Control-Allow-Origin") != originLocalhost {
		t.Errorf("Response should have 'Access-Control-Allow-Origin' header with value %s, but it's value was '%s'", originLocalhost, res.Header.Get("Access-Control-Allow-Origin"))
	}

	allowedHeaders := strings.Split(res.Header.Get("Access-Control-Allow-Headers"), ", ")
	if len(allowedHeaders) != 5 {
		t.Errorf("Response should have 'Access-Control-Allow-Headers' header with value %s, but it's value was '%s'", originLocalhost, res.Header.Get("Access-Control-Allow-Headers"))
	}

	allowedMethods := strings.Split(res.Header.Get("Access-Control-Allow-Methods"), ", ")
	if len(allowedMethods) != 2 {
		t.Errorf("Response should have 'Access-Control-Allow-Headers' header with value %s, but it's value was '%s'", originLocalhost, res.Header.Get("Access-Control-Allow-Methods"))
	}
}

func TestRequestNotInAllowedHosts(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler())
	s := httptest.NewServer(CORS(mux))
	defer s.Close()

	req, err := http.NewRequest(http.MethodGet, s.URL, &bytes.Reader{})
	if err != nil {
		t.Errorf("---\nFailed to create test request\n%s\n---", err.Error())
	}

	req.Header.Add("Origin", originNotAllowed)
	c := http.Client{}

	res, err := c.Do(req)
	if err != nil {
		t.Errorf("---\nFailed to execute test request\n%s\n---", err.Error())
	}

	if res.Header.Get("Access-Control-Allow-Origin") != "" {
		t.Errorf("Response should have not 'Access-Control-Allow-Origin' header, but it has")
	}
}
