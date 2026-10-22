package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoSnapshotPolicyResponse Response Object
type UpdateAutoSnapshotPolicyResponse struct {
	SnapshotPolicy *GetSnapshotPolicy `json:"snapshot_policy,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateAutoSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutoSnapshotPolicyResponse", string(data)}, " ")
}
