package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPermissionRequest Request Object
type ShowPermissionRequest struct {

	// 共享资源权限的ID。
	PermissionId string `json:"permission_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源权限版本。
	PermissionVersion *int32 `json:"permission_version,omitempty"`
}

func (o ShowPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPermissionRequest struct{}"
	}

	return strings.Join([]string{"ShowPermissionRequest", string(data)}, " ")
}
