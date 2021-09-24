package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstanceRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`

	Body *OpenGaussInstanceRequest `json:"body,omitempty"`
}

func (o CreateInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}
