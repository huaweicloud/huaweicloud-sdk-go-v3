package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageParticipant 与会者列表。
type PageParticipant struct {

	// 被邀请的与会者信息。包含预约会议时邀请的与会者和会中主持人邀请的与会者。
	Data *[]ParticipantInfo `json:"data,omitempty"`

	// 查询偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 每页的记录数。
	Limit *int32 `json:"limit,omitempty"`

	// 总记录数。
	Count *int32 `json:"count,omitempty"`
}

func (o PageParticipant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageParticipant struct{}"
	}

	return strings.Join([]string{"PageParticipant", string(data)}, " ")
}
