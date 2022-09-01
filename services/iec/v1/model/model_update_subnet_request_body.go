package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新子网的请求体。
type UpdateSubnetRequestBody struct {
	Subnet *UpdateSubnetOption `json:"subnet,omitempty" xml:"subnet"`
}

func (o UpdateSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSubnetRequestBody", string(data)}, " ")
}
