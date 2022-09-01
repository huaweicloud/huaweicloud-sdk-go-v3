package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchQosHistoryMeetingsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 查询条目数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// QoS会议列表，按照会议开始时间降序排序。
	Data           *[]QosConferenceInfo `json:"data,omitempty" xml:"data"`
	HttpStatusCode int                  `json:"-"`
}

func (o SearchQosHistoryMeetingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosHistoryMeetingsResponse struct{}"
	}

	return strings.Join([]string{"SearchQosHistoryMeetingsResponse", string(data)}, " ")
}
