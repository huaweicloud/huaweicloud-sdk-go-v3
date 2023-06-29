package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotStatisticsResponse Response Object
type ListSnapshotStatisticsResponse struct {

	// 快照统计信息。
	Statistics     *[]SnapshotsStatistic `json:"statistics,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListSnapshotStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotStatisticsResponse", string(data)}, " ")
}
