package model

import (
	"encoding/json"

	"strings"
)

// 更新VPC请求体
type UpdateVpcRequestBody struct {
	Vpc *UpdateVpcOption `json:"vpc,omitempty"`
}

func (o UpdateVpcRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVpcRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpcRequestBody", string(data)}, " ")
}
