package task

import "encoding/json"

type WithRawMessage struct {
	Data json.RawMessage `json:"data"`
}

type WithAny struct {
	Data any `json:"data"`
}

type CustomResourceSpec struct {
	Name string `json:"name"`
}
