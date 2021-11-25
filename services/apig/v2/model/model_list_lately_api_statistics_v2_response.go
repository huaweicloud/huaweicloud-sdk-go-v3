package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLatelyApiStatisticsV2Response struct {
	// 响应码

	Code *string `json:"code,omitempty"`
	// 返回消息

	Msg *string `json:"msg,omitempty"`
	// 开始时间的UTC的时间戳

	StartTime *int64 `json:"start_time,omitempty"`
	// 截止时间的UTC的时间戳

	EndTime *int64 `json:"end_time,omitempty"`
	// 统计指标的数据结构结构体

	List           *[]StatisticsApi `json:"list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListLatelyApiStatisticsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLatelyApiStatisticsV2Response struct{}"
	}

	return strings.Join([]string{"ListLatelyApiStatisticsV2Response", string(data)}, " ")
}
