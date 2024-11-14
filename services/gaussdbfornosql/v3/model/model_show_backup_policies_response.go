package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupPoliciesResponse Response Object
type ShowBackupPoliciesResponse struct {
	BackupPolicy   *ShowBackupPolicyResult `json:"backup_policy,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowBackupPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupPoliciesResponse", string(data)}, " ")
}
