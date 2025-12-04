package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnStatisticsResponse Response Object
type ListCdnStatisticsResponse struct {

	// 统计起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 采样时间间隔
	Interval *int32 `json:"interval,omitempty"`

	// cdn数据查询结果
	Result         map[string][]int64 `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCdnStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListCdnStatisticsResponse", string(data)}, " ")
}
