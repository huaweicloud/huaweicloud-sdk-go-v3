package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchStatisticUserInfoResponse struct {
	// 总记录数。

	Count *int32 `json:"count,omitempty"`
	// 查询条目数量。

	Limit *int32 `json:"limit,omitempty"`
	// 查询偏移量。

	Offset *int32 `json:"offset,omitempty"`
	// 会议用户数据按时间点统计的查询结果数组。

	Data           *[]StatisticUserDataItem `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o SearchStatisticUserInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchStatisticUserInfoResponse struct{}"
	}

	return strings.Join([]string{"SearchStatisticUserInfoResponse", string(data)}, " ")
}
