package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AudioModerationResultResultSegments struct {

	// 命中的风险片段
	Segment *string `json:"segment,omitempty" xml:"segment"`
}

func (o AudioModerationResultResultSegments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioModerationResultResultSegments struct{}"
	}

	return strings.Join([]string{"AudioModerationResultResultSegments", string(data)}, " ")
}
