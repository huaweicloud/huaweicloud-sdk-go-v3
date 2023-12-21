package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubnetRequestBody 子网请求体。
type CreateSubnetRequestBody struct {
	Subnet *CreateSubnetOption `json:"subnet"`
}

func (o CreateSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubnetRequestBody", string(data)}, " ")
}
