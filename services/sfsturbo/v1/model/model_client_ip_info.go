package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClientIpInfo 获取已挂载的客户端ip信息
type ClientIpInfo struct {

	// 获取已挂载的客户端ip信息
	Ips *string `json:"ips,omitempty"`
}

func (o ClientIpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientIpInfo struct{}"
	}

	return strings.Join([]string{"ClientIpInfo", string(data)}, " ")
}
