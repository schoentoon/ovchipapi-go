package ovchipapi

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

type loginResponse struct {
	Scope string `json:"scope"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IDToken string `json:"id_token"`
	AccessToken string `json:"access_token"`
}

// Login a user using the username and password. An authorization token will be returned.
func Login(username, password string) (string, error) {
	resp, err := postAndBody(loginUrl, url.Values{
		"scope":         {"openid"},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"grant_type":    {"password"},
		"username":      {username},
		"password":      {password},
	})
	if err != nil {
		return "", err
	}

	object := &loginResponse{}

	err = json.Unmarshal(resp, &object)
	if err != nil {
		return "", err
	}

	if strings.TrimSpace(object.IDToken) == "" {
		return "", errors.New("ovchipapi: Missing id_token")
	}

	return authorize(object.IDToken)
}

func authorize(idToken string) (string, error) {
	var authorizationToken string

	err := postAndResponse(authorizeUrl, url.Values{
		"authenticationToken": {idToken},
	}, &authorizationToken)

	return authorizationToken, err
}
