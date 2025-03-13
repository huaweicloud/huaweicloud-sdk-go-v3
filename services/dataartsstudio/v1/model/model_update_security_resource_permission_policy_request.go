package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityResourcePermissionPolicyRequest Request Object
type UpdateSecurityResourcePermissionPolicyRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 空间资源权限策略id。
	PolicyId string `json:"policy_id"`

	Body *PermissionResourcePolicyCreateDto `json:"body,omitempty"`
}

func (o UpdateSecurityResourcePermissionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityResourcePermissionPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityResourcePermissionPolicyRequest", string(data)}, " ")
}
