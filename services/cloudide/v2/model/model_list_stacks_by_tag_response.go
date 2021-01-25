package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListStacksByTagResponse struct {
	Stack *StacksTag `json:"stack,omitempty"`
	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListStacksByTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListStacksByTagResponse struct{}"
	}

	return strings.Join([]string{"ListStacksByTagResponse", string(data)}, " ")
}
