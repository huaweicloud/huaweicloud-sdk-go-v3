package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSnapshotPolicyRequestBody struct {
	SnapshotPolicy *UpdateSnapshotPolicy `json:"snapshot_policy,omitempty"`
}

func (o UpdateSnapshotPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotPolicyRequestBody", string(data)}, " ")
}
