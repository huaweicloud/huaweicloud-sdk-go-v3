package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateEipRequestBody 云堡垒机绑定\\解绑弹性公网IP请求对象。
type OperateEipRequestBody struct {

	// 弹性公网IP的ID,使用UUID格式。
	PublicipId string `json:"publicip_id"`

	// 弹性公网IP，使用IP地址。
	PublicEip string `json:"public_eip"`
}

func (o OperateEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateEipRequestBody struct{}"
	}

	return strings.Join([]string{"OperateEipRequestBody", string(data)}, " ")
}
