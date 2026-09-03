package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWdrSnapshotResponse Response Object
type ShowWdrSnapshotResponse struct {

	// WDR快照文件列表
	SnapshotList   *[]WdrSnapshot `json:"snapshot_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowWdrSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWdrSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ShowWdrSnapshotResponse", string(data)}, " ")
}
