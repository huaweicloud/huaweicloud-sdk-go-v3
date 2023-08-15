package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstancePeriodRequestNics struct {

	// 网卡对应的子网ID
	SubnetId string `json:"subnet_id"`

	// IP地址，不填或空字符串为自动分配
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o CreateInstancePeriodRequestNics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancePeriodRequestNics struct{}"
	}

	return strings.Join([]string{"CreateInstancePeriodRequestNics", string(data)}, " ")
}
