package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityDynamicMaskingPolicyRequest Request Object
type CreateSecurityDynamicMaskingPolicyRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DynamicMaskingPolicyCreateDto `json:"body,omitempty"`
}

func (o CreateSecurityDynamicMaskingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityDynamicMaskingPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityDynamicMaskingPolicyRequest", string(data)}, " ")
}
