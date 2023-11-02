package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotCrossRegionResponse Response Object
type ListSnapshotCrossRegionResponse struct {

	// 区域列表
	Regions *[]SnapshotRegion `json:"regions,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSnapshotCrossRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotCrossRegionResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotCrossRegionResponse", string(data)}, " ")
}
