package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicIp 云堡垒机实例弹性公网IP信息。可填写null值
type PublicIp struct {

	// 弹性IP和EIP只能有一个，弹性IP或EIP的ID。
	Id *string `json:"id,omitempty"`

	// 弹性IP地址。
	PublicEip *string `json:"public_eip,omitempty"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}
