package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityResourcePermissionPolicyRequest Request Object
type CreateSecurityResourcePermissionPolicyRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *PermissionResourcePolicyCreateDto `json:"body,omitempty"`
}

func (o CreateSecurityResourcePermissionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityResourcePermissionPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityResourcePermissionPolicyRequest", string(data)}, " ")
}
