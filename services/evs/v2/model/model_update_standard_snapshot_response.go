package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStandardSnapshotResponse Response Object
type UpdateStandardSnapshotResponse struct {
	Snapshot       *SnapshotDetailV5 `json:"snapshot,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateStandardSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStandardSnapshotResponse struct{}"
	}

	return strings.Join([]string{"UpdateStandardSnapshotResponse", string(data)}, " ")
}
