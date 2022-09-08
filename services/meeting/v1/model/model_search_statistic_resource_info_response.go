package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchStatisticResourceInfoResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 查询条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 会议已购资源使用数据按时间点统计的查询结果数组。
	Data           *[]StatisticResourceDataItem `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o SearchStatisticResourceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchStatisticResourceInfoResponse struct{}"
	}

	return strings.Join([]string{"SearchStatisticResourceInfoResponse", string(data)}, " ")
}
