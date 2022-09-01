package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AudioModerationResultResultDetails struct {

	// 音频片段开始时间
	StartTime *int32 `json:"start_time,omitempty" xml:"start_time"`

	// 音频片段审核处理建议： block：包含敏感信息，不通过 review：需要人工复检
	Suggestion *AudioModerationResultResultDetailsSuggestion `json:"suggestion,omitempty" xml:"suggestion"`

	// 音频片段结束时间
	EndTime *int32 `json:"end_time,omitempty" xml:"end_time"`

	// 音频片段标签
	Label *string `json:"label,omitempty" xml:"label"`

	// 音频片段文本内容
	AudioText *string `json:"audio_text,omitempty" xml:"audio_text"`

	// 命中的风险片段信息列表，如果命中语义算法模型，则该字段不会存在
	Segments *[]AudioModerationResultResultSegments `json:"segments,omitempty" xml:"segments"`
}

func (o AudioModerationResultResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioModerationResultResultDetails struct{}"
	}

	return strings.Join([]string{"AudioModerationResultResultDetails", string(data)}, " ")
}

type AudioModerationResultResultDetailsSuggestion struct {
	value string
}

type AudioModerationResultResultDetailsSuggestionEnum struct {
	BLOCK  AudioModerationResultResultDetailsSuggestion
	REVIEW AudioModerationResultResultDetailsSuggestion
}

func GetAudioModerationResultResultDetailsSuggestionEnum() AudioModerationResultResultDetailsSuggestionEnum {
	return AudioModerationResultResultDetailsSuggestionEnum{
		BLOCK: AudioModerationResultResultDetailsSuggestion{
			value: "block",
		},
		REVIEW: AudioModerationResultResultDetailsSuggestion{
			value: "review",
		},
	}
}

func (c AudioModerationResultResultDetailsSuggestion) Value() string {
	return c.value
}

func (c AudioModerationResultResultDetailsSuggestion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioModerationResultResultDetailsSuggestion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
