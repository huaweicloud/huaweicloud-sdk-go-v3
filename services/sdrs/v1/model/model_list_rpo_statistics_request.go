package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRpoStatisticsRequest struct {
	// 每次请求返回结果个数限制，取值范围为[0,1000]的正整数，默认值为1000。

	Limit *int32 `json:"limit,omitempty"`
	// 每次请求开始的下标，即偏移量，默认值为0。offset必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
	// 开始时间。默认格式为：\"yyyy-MM-dd HH:mm:ss.SSS\"，例：\"2019-04-01 12:00:00.000\"。

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间。默认格式为：\"yyyy-MM-dd HH:mm:ss.SSS\"，例：\"2019-04-01 12:00:00.000\"。

	EndTime *string `json:"end_time,omitempty"`
	// 资源类型。replication：表示查询复制对的RPO超标趋势记录。

	ResourceType *string `json:"resource_type,omitempty"`
}

func (o ListRpoStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRpoStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListRpoStatisticsRequest", string(data)}, " ")
}
