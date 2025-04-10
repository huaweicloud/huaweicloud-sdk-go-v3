package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroupsDetail IP地址组中包含的IP或网段列表。
type IpGroupsDetail struct {

	// IP地址或网段。支持IPv4。
	Ip string `json:"ip"`

	// 备注信息，最长255字符。
	Description string `json:"description"`
}

func (o IpGroupsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupsDetail struct{}"
	}

	return strings.Join([]string{"IpGroupsDetail", string(data)}, " ")
}
