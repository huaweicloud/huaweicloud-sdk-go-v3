package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListChartsRequest struct {
	// 日志组ID

	LogGroupId string `json:"log_group_id"`
	// 日志流ID

	LogStreamId string `json:"log_stream_id"`
	// 查询游标，初始传入0，后续从上一次的返回值中获取

	Offset *int32 `json:"offset,omitempty"`
	// 每页数据量，最大值为100

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListChartsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChartsRequest struct{}"
	}

	return strings.Join([]string{"ListChartsRequest", string(data)}, " ")
}
