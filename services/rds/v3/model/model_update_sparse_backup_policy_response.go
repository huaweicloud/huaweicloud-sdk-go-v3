package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSparseBackupPolicyResponse Response Object
type UpdateSparseBackupPolicyResponse struct {
	Body *interface{} `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSparseBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSparseBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateSparseBackupPolicyResponse", string(data)}, " ")
}
