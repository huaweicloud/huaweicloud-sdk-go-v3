package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupDownloadPolicyRequest Request Object
type UpdateBackupDownloadPolicyRequest struct {
	Body *UpdateBackupDownloadPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateBackupDownloadPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupDownloadPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateBackupDownloadPolicyRequest", string(data)}, " ")
}
