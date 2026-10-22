package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSnapshotPolicyResponse Response Object
type ShowAutoSnapshotPolicyResponse struct {
	SnapshotPolicy *GetSnapshotPolicy `json:"snapshot_policy,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowAutoSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoSnapshotPolicyResponse", string(data)}, " ")
}
