package auth

import (
	"errors"
	"net/http"
	"strings"
)

// header example is :
//
//	Authorization = ApiKey {Insert apikey here}
func GetApiKey(header http.Header) (string, error) {
	apiKey := header.Get("Authorization")

	if apiKey == "" {
		return "", errors.New("couldn't find the apikey")
	}

	val := strings.Split(apiKey, " ")

	if len(val) != 2 {
		return "", errors.New("error formating the apiKey")
	}
	if val[0] != "ApiKey" {
		return "", errors.New("error formating the apiKey")
	}

	return val[1], nil

}
