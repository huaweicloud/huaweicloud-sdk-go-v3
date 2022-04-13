package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotResp struct {
	Id *string `json:"id,omitempty"`
}

func (o SnapshotResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotResp struct{}"
	}

	return strings.Join([]string{"SnapshotResp", string(data)}, " ")
}
