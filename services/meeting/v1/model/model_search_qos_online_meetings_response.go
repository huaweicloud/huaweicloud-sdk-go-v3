package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchQosOnlineMeetingsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 查询条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// QoS会议列表，按照会议开始时间降序排序。
	Data           *[]QosConferenceInfo `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o SearchQosOnlineMeetingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosOnlineMeetingsResponse struct{}"
	}

	return strings.Join([]string{"SearchQosOnlineMeetingsResponse", string(data)}, " ")
}
