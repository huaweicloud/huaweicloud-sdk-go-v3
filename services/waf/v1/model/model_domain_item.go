package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 受攻击域名统计
type DomainItem struct {

	// 域名
	Key *string `json:"key,omitempty"`

	// 数量
	Num *int32 `json:"num,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`
}

func (o DomainItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainItem struct{}"
	}

	return strings.Join([]string{"DomainItem", string(data)}, " ")
}
