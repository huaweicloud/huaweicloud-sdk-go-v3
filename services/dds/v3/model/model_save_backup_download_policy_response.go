package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveBackupDownloadPolicyResponse Response Object
type SaveBackupDownloadPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveBackupDownloadPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveBackupDownloadPolicyResponse struct{}"
	}

	return strings.Join([]string{"SaveBackupDownloadPolicyResponse", string(data)}, " ")
}
