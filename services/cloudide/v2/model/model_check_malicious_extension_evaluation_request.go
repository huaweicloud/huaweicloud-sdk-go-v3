package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckMaliciousExtensionEvaluationRequest struct {
	Body *EvaluationAccusation `json:"body,omitempty"`
}

func (o CheckMaliciousExtensionEvaluationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMaliciousExtensionEvaluationRequest struct{}"
	}

	return strings.Join([]string{"CheckMaliciousExtensionEvaluationRequest", string(data)}, " ")
}
