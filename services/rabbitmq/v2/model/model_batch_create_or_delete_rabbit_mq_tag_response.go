package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateOrDeleteRabbitMqTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteRabbitMqTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteRabbitMqTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteRabbitMqTagResponse", string(data)}, " ")
}
