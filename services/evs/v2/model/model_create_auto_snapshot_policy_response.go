package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutoSnapshotPolicyResponse Response Object
type CreateAutoSnapshotPolicyResponse struct {
	SnapshotPolicy *GetSnapshotPolicy `json:"snapshot_policy,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateAutoSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateAutoSnapshotPolicyResponse", string(data)}, " ")
}
