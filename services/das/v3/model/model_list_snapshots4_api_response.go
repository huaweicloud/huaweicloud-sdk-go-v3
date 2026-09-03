package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshots4ApiResponse Response Object
type ListSnapshots4ApiResponse struct {

	// 锁快照信息
	Items *[]Snapshot `json:"items,omitempty"`

	// 快照总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSnapshots4ApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshots4ApiResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshots4ApiResponse", string(data)}, " ")
}
