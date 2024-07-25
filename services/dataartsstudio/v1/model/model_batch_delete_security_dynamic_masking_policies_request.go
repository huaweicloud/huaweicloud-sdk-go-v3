package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityDynamicMaskingPoliciesRequest Request Object
type BatchDeleteSecurityDynamicMaskingPoliciesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BatchDeletePolicySetsDto `json:"body,omitempty"`
}

func (o BatchDeleteSecurityDynamicMaskingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityDynamicMaskingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityDynamicMaskingPoliciesRequest", string(data)}, " ")
}
