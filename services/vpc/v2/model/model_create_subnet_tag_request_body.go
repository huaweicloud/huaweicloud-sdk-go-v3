package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateSubnetTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateSubnetTagRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSubnetTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubnetTagRequestBody", string(data)}, " ")
}
