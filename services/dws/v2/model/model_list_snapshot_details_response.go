package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSnapshotDetailsResponse struct {
	Snapshot       *SnapshoDetail `json:"snapshot,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSnapshotDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotDetailsResponse", string(data)}, " ")
}
