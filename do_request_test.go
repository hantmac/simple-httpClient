package simpleHttpClient

import (
	"github.com/hantmac/simple-httpClient"
	"testing"
)

func Test_DoRequest(t *testing.T) {
	headers := map[string]string{
		"Content-type": "application/json",
	}
	_, _, err := simpleHttpClient.DoRequest("GET", "https://api.github.com/events", headers, nil, 10)

	if err != nil {
		t.Error("do_request test not pass")
	} else {
		t.Log("do_request test passed")
	}

}
