package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 路由表
type RouteTable struct {

	// 路由表的id
	Id *string `json:"id,omitempty"`

	// 路由表名字
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 是否为默认关联的路由表
	IsDefaultAssociation bool `json:"is_default_association"`

	// 是否为默认传递路由表
	IsDefaultPropagation bool `json:"is_default_propagation"`

	// 路由表状态，支持的状态有pending | available | deleting | deleted | failed
	State string `json:"state"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o RouteTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteTable struct{}"
	}

	return strings.Join([]string{"RouteTable", string(data)}, " ")
}
