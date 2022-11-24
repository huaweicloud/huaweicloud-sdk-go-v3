package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSnapshotDetailsResponse struct {
	Snapshot       *SnapshotDetail `json:"snapshot,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSnapshotDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotDetailsResponse", string(data)}, " ")
}
