package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Permission 有关RAM权限的信息。
type Permission struct {

	// 权限ID。
	Id string `json:"id"`

	// 权限名称。
	Name string `json:"name"`

	// 资源类型。
	ResourceType string `json:"resource_type"`

	// 权限的影响和行为。
	Content string `json:"content"`

	// 该权限是否是此资源类型的默认权限。
	IsResourceTypeDefault bool `json:"is_resource_type_default"`

	// 权限的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 最后一次更新权限的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 权限URN。
	PermissionUrn *string `json:"permission_urn,omitempty"`

	// 权限类型。
	PermissionType *string `json:"permission_type,omitempty"`

	// 此权限是否为默认版本。
	DefaultVersion *bool `json:"default_version,omitempty"`

	// 权限版本。
	Version *int32 `json:"version,omitempty"`

	// 权限的状态
	Status *string `json:"status,omitempty"`
}

func (o Permission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Permission struct{}"
	}

	return strings.Join([]string{"Permission", string(data)}, " ")
}
