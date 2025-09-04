package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupDownloadPolicyResponse Response Object
type UpdateBackupDownloadPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBackupDownloadPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupDownloadPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateBackupDownloadPolicyResponse", string(data)}, " ")
}
