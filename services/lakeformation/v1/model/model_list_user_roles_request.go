package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserRolesRequest Request Object
type ListUserRolesRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// 模糊匹配角色名称。只能包含字母、数字和_|*.特殊字符，且长度为1~255个字符。
	RolePattern *string `json:"role_pattern,omitempty"`

	// 查询返回条数。默认值为100。最小值为1，最大值为1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID。最小长度为0，最大长度为1024。
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页。默认为false。
	ReversePage *bool `json:"reverse_page,omitempty"`

	// 用户名。只能包含字母、数字、下划线和中划线，且长度为1~256个字符。
	UserName string `json:"user_name"`
}

func (o ListUserRolesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserRolesRequest struct{}"
	}

	return strings.Join([]string{"ListUserRolesRequest", string(data)}, " ")
}
