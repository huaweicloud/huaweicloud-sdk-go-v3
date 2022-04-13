package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type WhitelistResp struct {
	// 白名单id

	Id string `json:"id"`
	// 白名单所在的项目ID

	TenantId string `json:"tenant_id"`
	// 白名单关联的监听器ID

	ListenerId string `json:"listener_id"`
	// 是否开启白名单访问控制开关。true：开启；false：关闭

	EnableWhitelist bool `json:"enable_whitelist"`
	// 白名单IP列表。可以是ip，例如192.168.10.123；也可以是一个网段，例如192.168.10.1/24；不同的值之间用逗号分隔

	Whitelist string `json:"whitelist"`
}

func (o WhitelistResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhitelistResp struct{}"
	}

	return strings.Join([]string{"WhitelistResp", string(data)}, " ")
}
