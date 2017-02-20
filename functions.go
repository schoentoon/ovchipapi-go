package ovchipapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// POST the values to the URl and return the body
func postAndBody(url string, values url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, values)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// POST the values to the URL and return the JSON decoded value
func postAndJson(url string, values url.Values) (map[string]*json.RawMessage, error) {
	body, err := postAndBody(url, values)
	if err != nil {
		return nil, err
	}

	var response map[string]*json.RawMessage
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// POST the values to the URL and return the "o" key deserialized into v
func postAndResponse(url string, values url.Values, v interface{}) error {
	resp, err := postAndJson(url, values)
	if err != nil {
		return err
	}

	var responseCode int
	err = json.Unmarshal(*resp["c"], &responseCode)
	if err != nil {
		return err
	}

	if responseCode != 200 {
		var output string
		if resp["o"] != nil {
			err = json.Unmarshal(*resp["o"], &output)
			if err != nil {
				return err
			}
		} else if resp["e"] != nil {
			err = json.Unmarshal(*resp["e"], &output)
			if err != nil {
				return err
			}
		}

		return fmt.Errorf("ovchipapi: API did not return successful response: %s", output)
	}

	err = json.Unmarshal(*resp["o"], v)
	if err != nil {
		return err
	}

	return nil
}
