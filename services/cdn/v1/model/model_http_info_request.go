package model

import (
	"encoding/json"

	"strings"
)

type HttpInfoRequest struct {
	Https *HttpInfoRequestBody `json:"https"`
}

func (o HttpInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HttpInfoRequest struct{}"
	}

	return strings.Join([]string{"HttpInfoRequest", string(data)}, " ")
}
