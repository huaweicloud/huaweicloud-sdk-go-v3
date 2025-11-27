package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecyclePolicyRequest Request Object
type UpdateRecyclePolicyRequest struct {
	Body *ModifyRecycleBinPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecyclePolicyRequest", string(data)}, " ")
}
