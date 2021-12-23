package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监听器HTTP扩展头部对象。
type InsertHeader struct {
	// 负载均衡器弹性公网IP透传开关。

	XForwardedELBIP *bool `json:"X-Forwarded-ELB-IP,omitempty"`
	// X-Forwarded-Host设为true可以将客户请求头的第一个X-Forwarded-Host设置为请求头的Host带到后端云服务器。

	XForwardedHost *bool `json:"X-Forwarded-Host,omitempty"`
}

func (o InsertHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsertHeader struct{}"
	}

	return strings.Join([]string{"InsertHeader", string(data)}, " ")
}
