package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditMediaTaskInput struct {
	Input *ObsInfo `json:"input"`

	// 剪切开始时间
	TimelineStart *string `json:"timeline_start,omitempty"`

	// 剪切结束时间
	TimelineEnd *string `json:"timeline_end,omitempty"`
}

func (o EditMediaTaskInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditMediaTaskInput struct{}"
	}

	return strings.Join([]string{"EditMediaTaskInput", string(data)}, " ")
}
