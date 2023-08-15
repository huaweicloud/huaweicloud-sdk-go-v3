package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotStatisticsRequest Request Object
type ListSnapshotStatisticsRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`
}

func (o ListSnapshotStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotStatisticsRequest", string(data)}, " ")
}
