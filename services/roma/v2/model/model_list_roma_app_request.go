package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRomaAppRequest Request Object
type ListRomaAppRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 查询收藏的应用 - 未提供时，查询当前用户有权限的所有应用 - 为true时，获取收藏的应用 - 为false时，获取未被收藏的应用
	Favorite *bool `json:"favorite,omitempty"`

	// 获取拥有指定权限应用
	AuthRole *string `json:"auth_role,omitempty"`

	// 应用名称，模糊匹配
	Name *string `json:"name,omitempty"`

	// 查询有权限访问的应用 - 未提供时，查询当前用户有权限的所有应用 - 为true时，查询当前用户创建的应用 - 为false时，查询非当前用户创建的有权限的应用，比如其它人共享的应用
	Owner *bool `json:"owner,omitempty"`

	// 从当前调用者有权限的所有应用中过滤出指定用户名有权限的应用
	UserName *string `json:"user_name,omitempty"`
}

func (o ListRomaAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRomaAppRequest struct{}"
	}

	return strings.Join([]string{"ListRomaAppRequest", string(data)}, " ")
}
