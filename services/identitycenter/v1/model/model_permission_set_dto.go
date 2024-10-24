package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PermissionSetDto 包含权限集详细信息的对象
type PermissionSetDto struct {

	// 权限集的创建时间
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 权限集的描述
	Description *string `json:"description,omitempty"`

	// 权限集的名称
	Name *string `json:"name,omitempty"`

	// 权限集的唯一标识
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 用于在联合身份验证过程中重定向应用程序中的用户
	RelayState *string `json:"relay_state,omitempty"`

	// 应用程序用户会话在ISO-8601标准中有效的时间长度
	SessionDuration *string `json:"session_duration,omitempty"`

	// 权限集的统一资源名称（URN）
	PermissionUrn *string `json:"permission_urn,omitempty"`
}

func (o PermissionSetDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetDto struct{}"
	}

	return strings.Join([]string{"PermissionSetDto", string(data)}, " ")
}
