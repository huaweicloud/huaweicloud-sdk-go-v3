package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcRequestBody 更新VPC请求体
type UpdateVpcRequestBody struct {
	Vpc *UpdateVpcOption `json:"vpc,omitempty"`
}

func (o UpdateVpcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpcRequestBody", string(data)}, " ")
}
