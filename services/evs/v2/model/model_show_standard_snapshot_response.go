package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStandardSnapshotResponse Response Object
type ShowStandardSnapshotResponse struct {
	Snapshot       *SnapshotDetailV5 `json:"snapshot,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowStandardSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStandardSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ShowStandardSnapshotResponse", string(data)}, " ")
}
