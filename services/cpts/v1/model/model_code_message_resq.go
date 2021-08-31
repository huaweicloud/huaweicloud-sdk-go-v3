package model

import (
	"encoding/json"

	"strings"
)

type CodeMessageResq struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`
}

func (o CodeMessageResq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CodeMessageResq struct{}"
	}

	return strings.Join([]string{"CodeMessageResq", string(data)}, " ")
}
