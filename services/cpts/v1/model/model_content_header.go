package model

import (
	"encoding/json"

	"strings"
)

type ContentHeader struct {
	// key

	Key *string `json:"key,omitempty"`
	// value

	Value *string `json:"value,omitempty"`
}

func (o ContentHeader) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ContentHeader struct{}"
	}

	return strings.Join([]string{"ContentHeader", string(data)}, " ")
}
