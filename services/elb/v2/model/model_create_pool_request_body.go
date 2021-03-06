package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreatePoolRequestBody struct {
	Pool *CreatePoolReq `json:"pool"`
}

func (o CreatePoolRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePoolRequestBody", string(data)}, " ")
}
