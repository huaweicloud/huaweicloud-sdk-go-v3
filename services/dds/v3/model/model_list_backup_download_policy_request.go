package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupDownloadPolicyRequest Request Object
type ListBackupDownloadPolicyRequest struct {
}

func (o ListBackupDownloadPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupDownloadPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListBackupDownloadPolicyRequest", string(data)}, " ")
}
