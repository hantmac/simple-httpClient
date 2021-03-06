**simpleHttpClient help you do http request easily**
* usage:

    go get github.com/hantmac/simple-httpClient

- Use simpleHttpClient you have no need to concern about handling error response or check the status code.
- You can use the lib flow the next example

```
headers := map[string]string{
{"Content-type":     "application/json"},
}

resp,header,err := simpleHttpClient.DoRequest("GET", "https://api.github.com/events", headers, nil, 10)

```
- That means the action you use to do request is GET, and the URL is "https://api.github.com/events",the headers of the request is a map which is {"Content-type": "application/json"}, the body for this request is nil(If you need the request body,see the next step), and finally you can add a number represent the time that you want to wait, if timeout, report an error.
### Use body
- if your request body is not empty, you can add and configure your body like this:
    - create a request body struct:
``` 
type requestBody struct {
            ID string
            name string
        }
```
-  New a object of request body
```
cb := &requestBody{}
```

- Add paramters to body

```
cb.ID = "test"
cb.name = "jack"
```

- marshal struct to []byte

```
body, err := json.Marshal(cb)
  ```

- Finally you can use your request body like:

```
resp, err : = simpleHttpClient.DoRequest("POST", 'http://httpbin.org/post', headers, body, timeoutSeconds)

```
