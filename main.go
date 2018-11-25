package main

import (
	"fmt"
	"github.com/hantmac/simple-httpClient"
	"github.com/sirupsen/logrus"
)

func main()  {
	headers := map[string]string{
		"Content-type":"application/json",
	}
	resp, header, err := simpleHttpClient.DoRequest("GET","https://api.github.com/events", headers, nil, 10)
	if err != nil {
		logrus.WithError(err).Fatal("Get response from github failed")
	}

	fmt.Printf("The response of github/events is %s:",resp)
	fmt.Println(header)
}
