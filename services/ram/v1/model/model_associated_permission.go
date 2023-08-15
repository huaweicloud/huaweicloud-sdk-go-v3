package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociatedPermission 描述与资源共享关联的权限的对象。
type AssociatedPermission struct {

	// 权限的ID。
	PermissionId *string `json:"permission_id,omitempty"`

	// 共享资源权限的名称。
	PermissionName *string `json:"permission_name,omitempty"`

	// 权限适用的资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 权限的当前状态。
	Status *string `json:"status,omitempty"`

	// 创建权限的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 最后一次更新权限的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o AssociatedPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatedPermission struct{}"
	}

	return strings.Join([]string{"AssociatedPermission", string(data)}, " ")
}
