package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnUserGroup struct {

	// 用户组ID
	Id *string `json:"id,omitempty"`

	// 用户组名称
	Name *string `json:"name,omitempty"`

	// 用户组描述
	Description *string `json:"description,omitempty"`

	// 用户组类型
	Type *string `json:"type,omitempty"`

	// 用户数量
	UserNumber *int32 `json:"user_number,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o VpnUserGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnUserGroup struct{}"
	}

	return strings.Join([]string{"VpnUserGroup", string(data)}, " ")
}
