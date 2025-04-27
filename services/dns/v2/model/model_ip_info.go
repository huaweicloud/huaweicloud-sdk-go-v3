package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpInfo struct {

	// IP地址。
	Ip *string `json:"ip,omitempty"`
}

func (o IpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpInfo struct{}"
	}

	return strings.Join([]string{"IpInfo", string(data)}, " ")
}
