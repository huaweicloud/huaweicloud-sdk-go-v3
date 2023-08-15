package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchStatisticConferenceInfoResponse Response Object
type SearchStatisticConferenceInfoResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 查询条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 会议总体数据按时间点统计的查询结果数组。
	Data           *[]StatisticConferenceDataItem `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o SearchStatisticConferenceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchStatisticConferenceInfoResponse struct{}"
	}

	return strings.Join([]string{"SearchStatisticConferenceInfoResponse", string(data)}, " ")
}
