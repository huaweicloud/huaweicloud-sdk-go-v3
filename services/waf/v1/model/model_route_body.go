package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个路由线路详细信息
type RouteBody struct {

	// WAF集群的cname后缀
	Cname *string `json:"cname,omitempty"`

	// WAF集群名称
	Name *string `json:"name,omitempty"`

	// 防护域名源站服务器信息列表
	Servers *[]RouteServerBody `json:"servers,omitempty"`
}

func (o RouteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteBody struct{}"
	}

	return strings.Join([]string{"RouteBody", string(data)}, " ")
}
