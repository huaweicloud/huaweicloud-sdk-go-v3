package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStatisticsApiResponse struct {
	// 响应码

	Code *string `json:"code,omitempty"`
	// 开始时间的UTC的时间戳

	StartTime *int64 `json:"start_time,omitempty"`
	// 截止时间的UTC的时间戳

	EndTime *int64 `json:"end_time,omitempty"`
	// 统计指标的数据结构结构体

	List           *[]StatisticsApi `json:"list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListStatisticsApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsApiResponse struct{}"
	}

	return strings.Join([]string{"ListStatisticsApiResponse", string(data)}, " ")
}
