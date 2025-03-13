package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityResourcePermissionPoliciesRequest Request Object
type BatchDeleteSecurityResourcePermissionPoliciesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BatchDeleteResourcePolicyDto `json:"body,omitempty"`
}

func (o BatchDeleteSecurityResourcePermissionPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityResourcePermissionPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityResourcePermissionPoliciesRequest", string(data)}, " ")
}
