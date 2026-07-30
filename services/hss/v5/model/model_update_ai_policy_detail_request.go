package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiPolicyDetailRequest Request Object
type UpdateAiPolicyDetailRequest struct {
	Body *UpdateAiPolicyDetailRequestInfo `json:"body,omitempty"`
}

func (o UpdateAiPolicyDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiPolicyDetailRequest struct{}"
	}

	return strings.Join([]string{"UpdateAiPolicyDetailRequest", string(data)}, " ")
}
