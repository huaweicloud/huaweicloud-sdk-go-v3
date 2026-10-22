package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAiGuardPolicyRequest Request Object
type BatchDeleteAiGuardPolicyRequest struct {
	Body *BatchDeleteAiGuardPolicyRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteAiGuardPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAiGuardPolicyRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAiGuardPolicyRequest", string(data)}, " ")
}
