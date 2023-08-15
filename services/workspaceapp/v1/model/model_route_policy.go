package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoutePolicy 路由策略
type RoutePolicy struct {

	// 单台服务器最大的链接会话数
	MaxSession *int32 `json:"max_session,omitempty"`
}

func (o RoutePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutePolicy struct{}"
	}

	return strings.Join([]string{"RoutePolicy", string(data)}, " ")
}
