package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchStatisticConferenceParticipantResponse Response Object
type SearchStatisticConferenceParticipantResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 查询条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 会议与会数据按时间点统计的查询结果数组。
	Data           *[]StatisticParticipateDataItem `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o SearchStatisticConferenceParticipantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchStatisticConferenceParticipantResponse struct{}"
	}

	return strings.Join([]string{"SearchStatisticConferenceParticipantResponse", string(data)}, " ")
}
