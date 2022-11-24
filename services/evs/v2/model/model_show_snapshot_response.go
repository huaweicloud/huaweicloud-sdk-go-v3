package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSnapshotResponse struct {
	Snapshot       *SnapshotDetails `json:"snapshot,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ShowSnapshotResponse", string(data)}, " ")
}
