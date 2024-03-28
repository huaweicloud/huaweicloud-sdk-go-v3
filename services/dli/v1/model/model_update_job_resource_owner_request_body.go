package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobResourceOwnerRequestBody group_name和resource_name可以单独使用，也可以组合使用 修改组的拥有者：使用group_name 修改资源包拥有者：使用resource_name 修改组下的资源包的拥有者：使用group_name加resource_name
type UpdateJobResourceOwnerRequestBody struct {

	// 新用户名,只能包含数字、英文字母、下划线和中划线且不能以数字开头,长度在5-32字符之间
	NewOwner string `json:"new_owner"`

	// 组名,名称只能包含数字、英文字母、点、下划线和中划线,长度不能超过64字符
	GroupName string `json:"group_name"`

	// 包名,包名,长度（包含文件后缀）不能超过128个字符
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o UpdateJobResourceOwnerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobResourceOwnerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateJobResourceOwnerRequestBody", string(data)}, " ")
}
