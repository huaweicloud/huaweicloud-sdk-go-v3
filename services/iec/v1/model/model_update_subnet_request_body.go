package model

import (
	"encoding/json"

	"strings"
)

// 更新子网的请求体。
type UpdateSubnetRequestBody struct {
	Subnet *UpdateSubnetOption `json:"subnet,omitempty"`
}

func (o UpdateSubnetRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSubnetRequestBody", string(data)}, " ")
}
