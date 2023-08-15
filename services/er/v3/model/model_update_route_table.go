package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRouteTable 更新路由表请求体
type UpdateRouteTable struct {

	// 路由器表名称
	Name *string `json:"name,omitempty"`

	// 路由器表描述信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateRouteTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteTable struct{}"
	}

	return strings.Join([]string{"UpdateRouteTable", string(data)}, " ")
}
