package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 网卡信息
type Nics struct {

	// 子网ID
	SubnetId string `json:"subnet_id"`

	// IP地址，不填或空字符串为自动分配
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o Nics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nics struct{}"
	}

	return strings.Join([]string{"Nics", string(data)}, " ")
}
