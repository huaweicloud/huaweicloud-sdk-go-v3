package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpaddressInfo struct {

	// 子网的ID。
	SubnetId string `json:"subnet_id"`

	// 自定义IP地址，需在子网的网段内部。
	Ip *string `json:"ip,omitempty"`
}

func (o IpaddressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpaddressInfo struct{}"
	}

	return strings.Join([]string{"IpaddressInfo", string(data)}, " ")
}
