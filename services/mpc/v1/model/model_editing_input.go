package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditingInput struct {
	Input *ObsObjInfo `json:"input"`

	// 剪切开始时间，单位：秒。
	TimelineStart *string `json:"timeline_start,omitempty"`

	// 剪切结束时间，单位：秒。
	TimelineEnd *string `json:"timeline_end,omitempty"`
}

func (o EditingInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditingInput struct{}"
	}

	return strings.Join([]string{"EditingInput", string(data)}, " ")
}
