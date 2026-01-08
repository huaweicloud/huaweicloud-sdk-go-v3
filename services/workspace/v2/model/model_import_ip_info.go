package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportIpInfo struct {

	// 编号。
	Number *string `json:"number,omitempty"`

	// ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 子网掩码。
	SubnetMask *string `json:"subnet_mask,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o ImportIpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIpInfo struct{}"
	}

	return strings.Join([]string{"ImportIpInfo", string(data)}, " ")
}
