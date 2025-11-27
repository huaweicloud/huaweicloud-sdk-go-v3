package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateBackupPolicyResponse Response Object
type AssociateBackupPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"AssociateBackupPolicyResponse", string(data)}, " ")
}
