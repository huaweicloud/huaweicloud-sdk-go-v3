package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRolesRequest Request Object
type ListRolesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 模糊匹配角色名称
	RolePattern *string `json:"role_pattern,omitempty"`

	// 返回的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页
	ReversePage *bool `json:"reverse_page,omitempty"`
}

func (o ListRolesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRolesRequest struct{}"
	}

	return strings.Join([]string{"ListRolesRequest", string(data)}, " ")
}
