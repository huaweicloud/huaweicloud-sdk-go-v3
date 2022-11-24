package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchQosParticipantsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 查询条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 会议与会者QoS告警列表。
	Data           *[]QosParticipantInfo `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o SearchQosParticipantsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosParticipantsResponse struct{}"
	}

	return strings.Join([]string{"SearchQosParticipantsResponse", string(data)}, " ")
}
