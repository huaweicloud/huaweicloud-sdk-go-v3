package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoute 创建路由请求体
type CreateRoute struct {

	// 路由目的地址
	Destination string `json:"destination"`

	// 路由下一跳指向的连接ID
	AttachmentId *string `json:"attachment_id,omitempty"`

	// 是否为黑洞路由，默认为false
	IsBlackhole *bool `json:"is_blackhole,omitempty"`
}

func (o CreateRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoute struct{}"
	}

	return strings.Join([]string{"CreateRoute", string(data)}, " ")
}
