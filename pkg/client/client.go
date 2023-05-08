package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Url string
}

func (c Client) MakePostRequest(endpoint string, body []byte) (string, error) {
	requestUrl := fmt.Sprintf("%s%s", c.Url, endpoint)
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		fmt.Println("Error making request")
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error ending request")
		return "", err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return "", err
	}

	return string(resBody), nil
}
