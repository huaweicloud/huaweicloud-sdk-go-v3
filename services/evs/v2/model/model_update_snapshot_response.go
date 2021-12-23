package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSnapshotResponse struct {
	Snapshot       *SnapshotDetails `json:"snapshot,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotResponse struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotResponse", string(data)}, " ")
}
