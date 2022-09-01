package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchQosParticipantsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 查询条目数量。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// QoS会议与会者列表。
	Data           *[]QosParticipantInfo `json:"data,omitempty" xml:"data"`
	HttpStatusCode int                   `json:"-"`
}

func (o SearchQosParticipantsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosParticipantsResponse struct{}"
	}

	return strings.Join([]string{"SearchQosParticipantsResponse", string(data)}, " ")
}
