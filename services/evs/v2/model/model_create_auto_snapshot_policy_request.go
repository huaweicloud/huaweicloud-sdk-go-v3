package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutoSnapshotPolicyRequest Request Object
type CreateAutoSnapshotPolicyRequest struct {
	Body *CreateSnapshotPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateAutoSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateAutoSnapshotPolicyRequest", string(data)}, " ")
}
