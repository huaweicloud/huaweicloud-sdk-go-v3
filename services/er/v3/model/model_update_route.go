package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoute 路由
type UpdateRoute struct {

	// 路由下一跳
	AttachmentId *string `json:"attachment_id,omitempty"`

	// 是否为黑洞路由
	IsBlackhole *bool `json:"is_blackhole,omitempty"`

	// 路由描述信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoute struct{}"
	}

	return strings.Join([]string{"UpdateRoute", string(data)}, " ")
}
