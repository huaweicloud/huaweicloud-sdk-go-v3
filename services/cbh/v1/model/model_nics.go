package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nics 云堡垒机实例网卡信息。
type Nics struct {

	// 子网ID，字母数字下划线连接符组成。
	SubnetId string `json:"subnet_id"`

	// IPV4地址。
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o Nics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nics struct{}"
	}

	return strings.Join([]string{"Nics", string(data)}, " ")
}
