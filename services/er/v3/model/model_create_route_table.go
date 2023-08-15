package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRouteTable 路由表
type CreateRouteTable struct {

	// 路由器表名称
	Name string `json:"name"`

	// 路由器表描述信息
	Description *string `json:"description,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateRouteTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTable struct{}"
	}

	return strings.Join([]string{"CreateRouteTable", string(data)}, " ")
}
