package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getBodyWithHeaders(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetJSONWithHeaders is a helper function to return a json object from a json rest api endpoint
func GetJSONWithHeaders(url string, headers map[string]string, v interface{}) error {
	body, err := getBodyWithHeaders(url, headers)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
