package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogsRequest Request Object
type ListSlowLogsRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0。取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty"`

	// 开始时间，UTC time，精确到毫秒。
	StartDate string `json:"start_date"`

	// 结束时间，UTC time，精确到毫秒。结束时间与开始时间，间隔不能超过7天。
	EndDate string `json:"end_date"`
}

func (o ListSlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogsRequest", string(data)}, " ")
}
