package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoModerationAudioDetailList struct {

	// 风险置信度
	Confidence *float32 `json:"confidence,omitempty"`

	// 风险标签
	Label *string `json:"label,omitempty"`

	// 审核处理建议： block：包含敏感信息，不通过 review：需要人工复检
	Suggestion *VideoModerationAudioDetailListSuggestion `json:"suggestion,omitempty"`

	// 命中的风险片段信息列表，如果命中语义算法模型，则该字段不会存在
	Segments *[]VideoModerationDetailSegment `json:"segments,omitempty"`
}

func (o VideoModerationAudioDetailList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationAudioDetailList struct{}"
	}

	return strings.Join([]string{"VideoModerationAudioDetailList", string(data)}, " ")
}

type VideoModerationAudioDetailListSuggestion struct {
	value string
}

type VideoModerationAudioDetailListSuggestionEnum struct {
	BLOCK  VideoModerationAudioDetailListSuggestion
	REVIEW VideoModerationAudioDetailListSuggestion
}

func GetVideoModerationAudioDetailListSuggestionEnum() VideoModerationAudioDetailListSuggestionEnum {
	return VideoModerationAudioDetailListSuggestionEnum{
		BLOCK: VideoModerationAudioDetailListSuggestion{
			value: "block",
		},
		REVIEW: VideoModerationAudioDetailListSuggestion{
			value: "review",
		},
	}
}

func (c VideoModerationAudioDetailListSuggestion) Value() string {
	return c.value
}

func (c VideoModerationAudioDetailListSuggestion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoModerationAudioDetailListSuggestion) UnmarshalJSON(b []byte) error {
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
