package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PermissionSummary 资源共享及其指定类型的任何资源关联的RAM权限的信息。
type PermissionSummary struct {

	// 权限ID。
	Id string `json:"id"`

	// 权限名称。
	Name string `json:"name"`

	// 此权限适用的资源类型。
	ResourceType string `json:"resource_type"`

	// 该权限是否是此资源类型的默认权限。
	IsResourceTypeDefault bool `json:"is_resource_type_default"`

	// 权限的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 上次更新权限的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o PermissionSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSummary struct{}"
	}

	return strings.Join([]string{"PermissionSummary", string(data)}, " ")
}
