package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnAccessPolicy struct {

	// 访问策略ID
	Id *string `json:"id,omitempty"`

	// 访问策略名称
	Name *string `json:"name,omitempty"`

	// 关联用户组ID
	UserGroupId *string `json:"user_group_id,omitempty"`

	// 关联用户组名称
	UserGroupName *string `json:"user_group_name,omitempty"`

	// 访问策略描述
	Description *string `json:"description,omitempty"`

	// 目的IP网段列表
	DestIpCidrs *[]string `json:"dest_ip_cidrs,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o VpnAccessPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnAccessPolicy struct{}"
	}

	return strings.Join([]string{"VpnAccessPolicy", string(data)}, " ")
}
