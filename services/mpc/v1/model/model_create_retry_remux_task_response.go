package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateRetryRemuxTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRetryRemuxTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRetryRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateRetryRemuxTaskResponse", string(data)}, " ")
}
