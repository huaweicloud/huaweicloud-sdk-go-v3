package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSnapshotCrossRegionPolicyResponse Response Object
type DeleteSnapshotCrossRegionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSnapshotCrossRegionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotCrossRegionPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotCrossRegionPolicyResponse", string(data)}, " ")
}
