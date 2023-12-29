package main

import (
	"encoding/json"
	"fmt"
	"github.com/dnguyenzd/task-schema-test/internal/task"
	"io"
	"net/http"
	"reflect"
	"time"
)

const host = "http://localhost:3000"

var endpoints = []string{
	"raw-message",
	"any",
}

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	for _, e := range endpoints {
		url := fmt.Sprintf("%s/%s", host, e)
		resp, err := httpClient.Get(url)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("endpoint /%s response body as string: %s\n", e, string(responseBody))

		var withRawMessage task.WithRawMessage
		if err := json.Unmarshal(responseBody, &withRawMessage); err != nil {
			panic(err)
		}
		fmt.Printf("endpoint /%s - unmarshalled into raw message, data type: %v, data: %+v\n", e, reflect.TypeOf(withRawMessage.Data), withRawMessage)

		var withAny task.WithAny
		if err := json.Unmarshal(responseBody, &withAny); err != nil {
			panic(err)
		}
		fmt.Printf("endpoint /%s - unmarshalled into any, data type: %v, data: %+v\n", e, reflect.TypeOf(withAny.Data), withAny)
	}
}
