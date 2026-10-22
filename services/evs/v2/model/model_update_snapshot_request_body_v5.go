package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSnapshotRequestBodyV5 struct {
	Snapshot *CinderUpdateSnapshotOptionV5 `json:"snapshot"`
}

func (o UpdateSnapshotRequestBodyV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotRequestBodyV5 struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotRequestBodyV5", string(data)}, " ")
}
