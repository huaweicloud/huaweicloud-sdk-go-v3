package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AudioModerationResultDetail struct {

	// 音频片段开始时间
	StartTime *float32 `json:"start_time,omitempty"`

	// 音频片段审核处理建议： block：包含敏感信息，不通过 review：需要人工复检
	Suggestion *AudioModerationResultDetailSuggestion `json:"suggestion,omitempty"`

	// 音频片段结束时间
	EndTime *float32 `json:"end_time,omitempty"`

	// 音频片段标签
	Label *string `json:"label,omitempty"`

	// 音频片段文本内容
	AudioText *string `json:"audio_text,omitempty"`

	// 命中的风险片段信息列表，如果命中语义算法模型，则该字段不会存在
	Segments *[]VideoModerationDetailSegment `json:"segments,omitempty"`
}

func (o AudioModerationResultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioModerationResultDetail struct{}"
	}

	return strings.Join([]string{"AudioModerationResultDetail", string(data)}, " ")
}

type AudioModerationResultDetailSuggestion struct {
	value string
}

type AudioModerationResultDetailSuggestionEnum struct {
	BLOCK  AudioModerationResultDetailSuggestion
	REVIEW AudioModerationResultDetailSuggestion
}

func GetAudioModerationResultDetailSuggestionEnum() AudioModerationResultDetailSuggestionEnum {
	return AudioModerationResultDetailSuggestionEnum{
		BLOCK: AudioModerationResultDetailSuggestion{
			value: "block",
		},
		REVIEW: AudioModerationResultDetailSuggestion{
			value: "review",
		},
	}
}

func (c AudioModerationResultDetailSuggestion) Value() string {
	return c.value
}

func (c AudioModerationResultDetailSuggestion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioModerationResultDetailSuggestion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
