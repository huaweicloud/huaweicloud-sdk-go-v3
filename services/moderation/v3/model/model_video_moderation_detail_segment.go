package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoModerationDetailSegment struct {

	// 命中的风险片段
	Segment *string `json:"segment,omitempty"`
}

func (o VideoModerationDetailSegment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationDetailSegment struct{}"
	}

	return strings.Join([]string{"VideoModerationDetailSegment", string(data)}, " ")
}
