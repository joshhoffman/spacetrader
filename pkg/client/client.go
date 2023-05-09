package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Url    string
	Token  string
	client http.Client
}

func (c Client) MakeGetRequest(endpoint string, responseObject MarshalableObject) (MarshalableObject, error) {
	requestUrl := fmt.Sprintf("%s%s", c.Url, endpoint)
	fmt.Println(requestUrl)

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		fmt.Println("Error making request")
		return nil, err
	}

	authHeader := fmt.Sprintf("Bearer %s", c.Token)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request %s\n", err)
		return nil, err
	}

	resBody, err := GetBody(res)
	if err != nil {
		fmt.Println("Error reading body")
		return nil, err
	}

	// println(string(resBody))

	responseObject, err = responseObject.Unmarshal(resBody)

	return responseObject, err
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

	res, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request %s\n", err)
		return "", err
	}

	resBody, err := GetBody(res)
	if err != nil {
		fmt.Println("Error reading body")
		return "", err
	}

	return string(resBody), nil
}

func GetBody(body *http.Response) ([]byte, error) {
	resBody, err := ioutil.ReadAll(body.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return nil, err
	}
	return resBody, nil
}

func MakeClient(url string, token string) Client {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return Client{Url: url, Token: token, client: client}
}
