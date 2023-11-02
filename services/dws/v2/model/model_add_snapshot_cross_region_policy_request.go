package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSnapshotCrossRegionPolicyRequest Request Object
type AddSnapshotCrossRegionPolicyRequest struct {
	Body *AddSnapshotCrossRegionPolicyRequestBody `json:"body,omitempty"`
}

func (o AddSnapshotCrossRegionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSnapshotCrossRegionPolicyRequest struct{}"
	}

	return strings.Join([]string{"AddSnapshotCrossRegionPolicyRequest", string(data)}, " ")
}
