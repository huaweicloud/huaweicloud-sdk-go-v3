package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRedisPitrRestoreTimeRequest Request Object
type ListRedisPitrRestoreTimeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 查询可恢复时间点的开始时间，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。
	StartTime string `json:"start_time"`

	// 查询可恢复时间点的结束时间，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。
	EndTime string `json:"end_time"`

	// 偏移量，表示查询该偏移量后面的记录，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量上限值，取值范围为1~100，默认值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRedisPitrRestoreTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedisPitrRestoreTimeRequest struct{}"
	}

	return strings.Join([]string{"ListRedisPitrRestoreTimeRequest", string(data)}, " ")
}
