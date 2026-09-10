package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionsRoute 用户添加的路由信息。
type ConnectionsRoute struct {

	// 路由名称，长度限制：1-64个字符。
	Name *string `json:"name,omitempty"`

	// 路由网段范围。
	Cidr *string `json:"cidr,omitempty"`

	// 创建路由时间。
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o ConnectionsRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionsRoute struct{}"
	}

	return strings.Join([]string{"ConnectionsRoute", string(data)}, " ")
}
