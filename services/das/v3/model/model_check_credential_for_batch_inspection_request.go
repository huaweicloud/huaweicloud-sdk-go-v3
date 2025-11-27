package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCredentialForBatchInspectionRequest Request Object
type CheckCredentialForBatchInspectionRequest struct {
	Body *CheckCredentialRequestBody `json:"body,omitempty"`
}

func (o CheckCredentialForBatchInspectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCredentialForBatchInspectionRequest struct{}"
	}

	return strings.Join([]string{"CheckCredentialForBatchInspectionRequest", string(data)}, " ")
}
