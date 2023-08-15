package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VideoModerationResultResult 作业审核结果，当作业状态为succeeded时存在
type VideoModerationResultResult struct {

	// 视频审核结果是否通过。 block：包含敏感信息，不通过  review：需要人工复检 pass：不包含敏感信息，通过
	Suggestion *VideoModerationResultResultSuggestion `json:"suggestion,omitempty"`

	// 图像审核详情
	ImageDetail *[]VideoModerationImageDetail `json:"image_detail,omitempty"`

	// 音频审核详情
	AudioDetail *[]VideoModerationVideoDetail `json:"audio_detail,omitempty"`
}

func (o VideoModerationResultResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationResultResult struct{}"
	}

	return strings.Join([]string{"VideoModerationResultResult", string(data)}, " ")
}

type VideoModerationResultResultSuggestion struct {
	value string
}

type VideoModerationResultResultSuggestionEnum struct {
	BLOCK  VideoModerationResultResultSuggestion
	PASS   VideoModerationResultResultSuggestion
	REVIEW VideoModerationResultResultSuggestion
}

func GetVideoModerationResultResultSuggestionEnum() VideoModerationResultResultSuggestionEnum {
	return VideoModerationResultResultSuggestionEnum{
		BLOCK: VideoModerationResultResultSuggestion{
			value: "block",
		},
		PASS: VideoModerationResultResultSuggestion{
			value: "pass",
		},
		REVIEW: VideoModerationResultResultSuggestion{
			value: "review",
		},
	}
}

func (c VideoModerationResultResultSuggestion) Value() string {
	return c.value
}

func (c VideoModerationResultResultSuggestion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoModerationResultResultSuggestion) UnmarshalJSON(b []byte) error {
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
