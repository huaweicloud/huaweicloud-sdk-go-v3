package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Ipaddresses struct {

	// 子网的网络id。
	SubnetId string `json:"subnet_id"`

	// 自定义ip地址，需在子网的网段内部。
	Ip *string `json:"ip,omitempty"`
}

func (o Ipaddresses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ipaddresses struct{}"
	}

	return strings.Join([]string{"Ipaddresses", string(data)}, " ")
}
