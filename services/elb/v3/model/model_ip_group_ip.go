package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroupIp IP地址组的IP地址对象。
type IpGroupIp struct {

	// 参数解释：IP地址，可以是具体的IP地址或者IP地址段。
	Ip string `json:"ip"`
}

func (o IpGroupIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupIp struct{}"
	}

	return strings.Join([]string{"IpGroupIp", string(data)}, " ")
}
