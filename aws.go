package goaws

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestParams struct {
	Auth   *Auth
	Url    string
	Method string
}

func GetCurrentDate() string {
	response, err := http.Get("https://route53.amazonaws.com/date")
	if err != nil {
		log.Print("Unable to get date from amazon: ", err)
		return ""
	}
	defer response.Body.Close()
	date := response.Header.Get("Date")
	return date
}

func Request(params *RequestParams) (string, error) {
	if params.Url == "" {
		return "", errors.New("No Url parameter given")
	}
	if params.Auth == nil {
		return "", errors.New("No Auth given")
	}

  date := GetCurrentDate()
	if date == "" {
		return "", errors.New("Unable to fetch Amazons reference date")
	}

	client := &http.Client{}

	method := "GET"
	if params.Method != "" {
		method = params.Method
	}

	req, err := http.NewRequest(method, params.Url, nil)
	requestHeader, err := params.Auth.GetHeader(date)
	if err != nil {
		return "", errors.New("Failed to create Authorization Headers")
	}
	req.Header.Add("X-Amzn-Authorization", requestHeader)
	req.Header.Add("X-Amz-Date", date)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
  body := string(respBody)

	return body, nil
}
