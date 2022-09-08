package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoModerationVideoDetail struct {

	// 音频片段审核结果是否通过。 block：包含敏感信息，不通过  review：需要人工复检
	Suggestion *VideoModerationVideoDetailSuggestion `json:"suggestion,omitempty"`

	// 音频片段检测标签，选取detail中置信度最大的标签，可取值如下： politics: 涉政  terrorism: 暴恐  porn: 色情  ad: 广告 ad_law: 广告法 abuse: 辱骂 ban: 违禁 meaningless: 无意义 moan: 娇喘
	Label *string `json:"label,omitempty"`

	// 音频片段文本内容
	AudioText *string `json:"audio_text,omitempty"`

	// 音频片段结束时间
	EndTime *int32 `json:"end_time,omitempty"`

	// 音频片段开始时间
	StartTime *int32 `json:"start_time,omitempty"`

	// 音频片段审核详情
	Detail *[]VideoModerationAudioDetailList `json:"detail,omitempty"`
}

func (o VideoModerationVideoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationVideoDetail struct{}"
	}

	return strings.Join([]string{"VideoModerationVideoDetail", string(data)}, " ")
}

type VideoModerationVideoDetailSuggestion struct {
	value string
}

type VideoModerationVideoDetailSuggestionEnum struct {
	BLOCK  VideoModerationVideoDetailSuggestion
	REVIEW VideoModerationVideoDetailSuggestion
}

func GetVideoModerationVideoDetailSuggestionEnum() VideoModerationVideoDetailSuggestionEnum {
	return VideoModerationVideoDetailSuggestionEnum{
		BLOCK: VideoModerationVideoDetailSuggestion{
			value: "block",
		},
		REVIEW: VideoModerationVideoDetailSuggestion{
			value: "review",
		},
	}
}

func (c VideoModerationVideoDetailSuggestion) Value() string {
	return c.value
}

func (c VideoModerationVideoDetailSuggestion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoModerationVideoDetailSuggestion) UnmarshalJSON(b []byte) error {
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
