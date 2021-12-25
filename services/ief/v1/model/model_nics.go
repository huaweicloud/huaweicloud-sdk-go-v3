package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Nics struct {
	// 边缘节点网卡名称

	Interface *string `json:"interface,omitempty"`
	// 上述网卡对应的IPv4地址

	Ip *string `json:"ip,omitempty"`
}

func (o Nics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nics struct{}"
	}

	return strings.Join([]string{"Nics", string(data)}, " ")
}
