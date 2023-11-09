package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DocumentAudioDetail struct {

	// 音频片段开始时间
	StartTime *float32 `json:"start_time,omitempty"`

	// 音频片段结束时间
	EndTime *float32 `json:"end_time,omitempty"`

	// 音频片段审核处理建议： block：包含敏感信息，不通过 review：需要人工复检
	Suggestion *string `json:"suggestion,omitempty"`

	// 音频片段文本内容
	AudioText *string `json:"audio_text,omitempty"`

	// 音频片段标签
	Label *string `json:"label,omitempty"`

	// 命中的风险片段信息列表
	Segments *[]DocumentVideoImageDetailSegments `json:"segments,omitempty"`
}

func (o DocumentAudioDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentAudioDetail struct{}"
	}

	return strings.Join([]string{"DocumentAudioDetail", string(data)}, " ")
}
