package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnUser struct {

	// 用户ID
	Id *string `json:"id,omitempty"`

	// 用户名称
	Name *string `json:"name,omitempty"`

	// 用户描述
	Description *string `json:"description,omitempty"`

	// 所属用户组ID
	UserGroupId *string `json:"user_group_id,omitempty"`

	// 所属用户组名称
	UserGroupName *string `json:"user_group_name,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o VpnUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnUser struct{}"
	}

	return strings.Join([]string{"VpnUser", string(data)}, " ")
}
