package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparseBackupPolicyResponse Response Object
type ListSparseBackupPolicyResponse struct {

	// 备份策略集
	Policies *[]SparseBackupPolicy `json:"policies,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSparseBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparseBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListSparseBackupPolicyResponse", string(data)}, " ")
}
