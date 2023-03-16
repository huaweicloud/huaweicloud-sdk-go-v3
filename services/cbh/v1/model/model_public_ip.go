package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云堡垒机实例弹性公网IP信息。
type PublicIp struct {

	// 已分配的弹性IP和EIP只能有一个。
	Id string `json:"id"`

	// 已分配的弹性IP的地址Address。
	PublicEip string `json:"public_eip"`

	Eip *Eip `json:"eip,omitempty"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}
