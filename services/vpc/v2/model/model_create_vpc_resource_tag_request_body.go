package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateVpcResourceTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateVpcResourceTagRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVpcResourceTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpcResourceTagRequestBody", string(data)}, " ")
}
