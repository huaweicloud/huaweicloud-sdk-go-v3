package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSnapshotPolicyResponse Response Object
type UpdateSnapshotPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotPolicyResponse", string(data)}, " ")
}
