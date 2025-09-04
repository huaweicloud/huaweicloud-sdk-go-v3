package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotsResponse Response Object
type ListSnapshotsResponse struct {

	// 具体的snapshot
	Items *[]QuerySnapshotsRespItems `json:"items,omitempty"`

	// snapshot的数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotsResponse", string(data)}, " ")
}
