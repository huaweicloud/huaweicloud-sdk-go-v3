package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveCredentialForBatchInspectionRequest Request Object
type SaveCredentialForBatchInspectionRequest struct {
	Body *SaveCredentialRequestBody `json:"body,omitempty"`
}

func (o SaveCredentialForBatchInspectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveCredentialForBatchInspectionRequest struct{}"
	}

	return strings.Join([]string{"SaveCredentialForBatchInspectionRequest", string(data)}, " ")
}
