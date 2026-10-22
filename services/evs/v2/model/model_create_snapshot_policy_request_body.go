package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSnapshotPolicyRequestBody struct {
	SnapshotPolicy *CreateSnapshotPolicy `json:"snapshot_policy"`
}

func (o CreateSnapshotPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSnapshotPolicyRequestBody", string(data)}, " ")
}
