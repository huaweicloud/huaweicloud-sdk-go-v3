package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UpdateWhitelistReq struct {
	// 是否开启白名单访问控制开关。true：开启；false：关闭

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`
	// 白名单IP列表。可以是ip，例如192.168.10.123；也可以是一个网段，例如192.168.10.1/24；不同的值之间用逗号分隔

	Whitelist *string `json:"whitelist,omitempty"`
}

func (o UpdateWhitelistReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhitelistReq struct{}"
	}

	return strings.Join([]string{"UpdateWhitelistReq", string(data)}, " ")
}
