package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateOrDeleteKafkaTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteKafkaTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteKafkaTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteKafkaTagResponse", string(data)}, " ")
}
