package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoModerationImageDetail struct {

	// 图像审核结果是否通过。 block：包含敏感信息，不通过  review：需要人工复检
	Suggestion *VideoModerationImageDetailSuggestion `json:"suggestion,omitempty"`

	// 检测结果的一级标签。 支持category列表如下： politics: 涉政  terrorism: 暴恐  porn: 色情  image_text: 图文审核
	Category *VideoModerationImageDetailCategory `json:"category,omitempty"`

	// 图文审核检测出的文本，只有在category参数配置image_text且检测出文本时展示该字段。
	OcrText *string `json:"ocr_text,omitempty"`

	// 截帧在视频文件中的时间，单位为秒
	Time *float32 `json:"time,omitempty"`

	// 图像帧审核详情
	Detail *[]VideoModerationImageDetailList `json:"detail,omitempty"`
}

func (o VideoModerationImageDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationImageDetail struct{}"
	}

	return strings.Join([]string{"VideoModerationImageDetail", string(data)}, " ")
}

type VideoModerationImageDetailSuggestion struct {
	value string
}

type VideoModerationImageDetailSuggestionEnum struct {
	BLOCK  VideoModerationImageDetailSuggestion
	REVIEW VideoModerationImageDetailSuggestion
}

func GetVideoModerationImageDetailSuggestionEnum() VideoModerationImageDetailSuggestionEnum {
	return VideoModerationImageDetailSuggestionEnum{
		BLOCK: VideoModerationImageDetailSuggestion{
			value: "block",
		},
		REVIEW: VideoModerationImageDetailSuggestion{
			value: "review",
		},
	}
}

func (c VideoModerationImageDetailSuggestion) Value() string {
	return c.value
}

func (c VideoModerationImageDetailSuggestion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoModerationImageDetailSuggestion) UnmarshalJSON(b []byte) error {
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

type VideoModerationImageDetailCategory struct {
	value string
}

type VideoModerationImageDetailCategoryEnum struct {
	POLITICS   VideoModerationImageDetailCategory
	TERRORISM  VideoModerationImageDetailCategory
	PORN       VideoModerationImageDetailCategory
	IMAGE_TEXT VideoModerationImageDetailCategory
}

func GetVideoModerationImageDetailCategoryEnum() VideoModerationImageDetailCategoryEnum {
	return VideoModerationImageDetailCategoryEnum{
		POLITICS: VideoModerationImageDetailCategory{
			value: "politics",
		},
		TERRORISM: VideoModerationImageDetailCategory{
			value: "terrorism",
		},
		PORN: VideoModerationImageDetailCategory{
			value: "porn",
		},
		IMAGE_TEXT: VideoModerationImageDetailCategory{
			value: "image_text",
		},
	}
}

func (c VideoModerationImageDetailCategory) Value() string {
	return c.value
}

func (c VideoModerationImageDetailCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoModerationImageDetailCategory) UnmarshalJSON(b []byte) error {
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
