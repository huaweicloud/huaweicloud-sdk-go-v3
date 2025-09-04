package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveBackupDownloadPolicyRequest Request Object
type SaveBackupDownloadPolicyRequest struct {
	Body *SaveBackupDownloadPolicyRequestBody `json:"body,omitempty"`
}

func (o SaveBackupDownloadPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveBackupDownloadPolicyRequest struct{}"
	}

	return strings.Join([]string{"SaveBackupDownloadPolicyRequest", string(data)}, " ")
}
