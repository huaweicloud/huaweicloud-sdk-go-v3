package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityDynamicMaskingPolicyRequest Request Object
type UpdateSecurityDynamicMaskingPolicyRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 脱敏策略id
	Id string `json:"id"`

	Body *DynamicMaskingPolicyUpdateDto `json:"body,omitempty"`
}

func (o UpdateSecurityDynamicMaskingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityDynamicMaskingPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityDynamicMaskingPolicyRequest", string(data)}, " ")
}
