package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSlowlogStatisticsResponse struct {

	// 当前页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber"`

	// 每页条数
	PageRecord *int32 `json:"pageRecord,omitempty" xml:"pageRecord"`

	// 慢日志列表
	SlowLogList *[]SlowLogStatistics `json:"slowLogList,omitempty" xml:"slowLogList"`

	// 总条数
	TotalRecord *int32 `json:"totalRecord,omitempty" xml:"totalRecord"`

	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime"`

	// 结束时间
	EndTime        *int64 `json:"endTime,omitempty" xml:"endTime"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowlogStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowlogStatisticsResponse", string(data)}, " ")
}
