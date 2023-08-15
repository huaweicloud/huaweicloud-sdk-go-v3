package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Route 路由表项
type Route struct {

	// 路由id
	Id string `json:"id"`

	// 路由类型
	Type string `json:"type"`

	// 路由状态
	State *string `json:"state,omitempty"`

	// 是否为黑洞路由
	IsBlackhole *bool `json:"is_blackhole,omitempty"`

	// 路由目的地址
	Destination string `json:"destination"`

	// 下一跳列表
	Attachments []RouteAttachment `json:"attachments"`

	// 路由表id
	RouteTableId string `json:"route_table_id"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o Route) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Route struct{}"
	}

	return strings.Join([]string{"Route", string(data)}, " ")
}
