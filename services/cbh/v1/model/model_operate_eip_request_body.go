package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 管理CBH实例Eip响应对象
type OperateEipRequestBody struct {

	// 弹性ip的id
	PublicipId string `json:"publicip_id"`

	// 弹性ip
	PublicEip string `json:"public_eip"`
}

func (o OperateEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateEipRequestBody struct{}"
	}

	return strings.Join([]string{"OperateEipRequestBody", string(data)}, " ")
}
