package model

import (
	"encoding/json"

	"strings"
)

type HeaderBody struct {
	Headers *HeaderMap `json:"headers,omitempty"`
}

func (o HeaderBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HeaderBody struct{}"
	}

	return strings.Join([]string{"HeaderBody", string(data)}, " ")
}
