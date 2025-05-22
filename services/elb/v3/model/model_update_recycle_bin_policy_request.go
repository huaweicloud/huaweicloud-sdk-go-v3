package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecycleBinPolicyRequest Request Object
type UpdateRecycleBinPolicyRequest struct {
	Body *UpdateRecycleBinPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateRecycleBinPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinPolicyRequest", string(data)}, " ")
}
