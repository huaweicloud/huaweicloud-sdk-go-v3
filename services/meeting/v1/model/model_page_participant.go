package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 与会者列表
type PageParticipant struct {

	// 与会者信息。
	Data *[]ParticipantInfo `json:"data,omitempty" xml:"data"`

	// 记录数偏移，这一页之前共有多少条。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页的记录数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 总记录数。
	Count *int32 `json:"count,omitempty" xml:"count"`
}

func (o PageParticipant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageParticipant struct{}"
	}

	return strings.Join([]string{"PageParticipant", string(data)}, " ")
}
