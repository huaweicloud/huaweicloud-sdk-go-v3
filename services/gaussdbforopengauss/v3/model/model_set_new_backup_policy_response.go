package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetNewBackupPolicyResponse Response Object
type SetNewBackupPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetNewBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetNewBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetNewBackupPolicyResponse", string(data)}, " ")
}
