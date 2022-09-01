package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotResp struct {

	// Snapshot ID
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o SnapshotResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotResp struct{}"
	}

	return strings.Join([]string{"SnapshotResp", string(data)}, " ")
}
