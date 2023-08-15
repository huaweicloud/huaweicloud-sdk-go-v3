package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoModerationImageDetailList struct {

	// 置信度，可选值在0-1之间，值越大，可信度越高。
	Confidence *float32 `json:"confidence,omitempty"`

	// 检测结果的一级标签。 支持category列表如下： politics: 涉政  terrorism: 暴恐  porn: 色情  image_text: 图文审核
	Category *VideoModerationImageDetailListCategory `json:"category,omitempty"`

	// 审核结果是否通过。  block：包含敏感信息，不通过  review：需要人工复检
	Suggestion *VideoModerationImageDetailListSuggestion `json:"suggestion,omitempty"`

	// 识别的详细标签。
	Label *string `json:"label,omitempty"`

	FaceLocation *VideoModerationImageDetailListFaceLocation `json:"face_location,omitempty"`

	QrLocation *VideoModerationImageDetailListQrLocation `json:"qr_location,omitempty"`

	// 图片中二维码指向的链接，当请求参数categories中包含image_text时存在。
	QrContent *string `json:"qr_content,omitempty"`

	// image_text场景下命中的文本片段。
	Segments *[]VideoModerationDetailSegment `json:"segments,omitempty"`
}

func (o VideoModerationImageDetailList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationImageDetailList struct{}"
	}

	return strings.Join([]string{"VideoModerationImageDetailList", string(data)}, " ")
}

type VideoModerationImageDetailListCategory struct {
	value string
}

type VideoModerationImageDetailListCategoryEnum struct {
	POLITICS   VideoModerationImageDetailListCategory
	TERRORISM  VideoModerationImageDetailListCategory
	PORN       VideoModerationImageDetailListCategory
	IMAGE_TEXT VideoModerationImageDetailListCategory
}

func GetVideoModerationImageDetailListCategoryEnum() VideoModerationImageDetailListCategoryEnum {
	return VideoModerationImageDetailListCategoryEnum{
		POLITICS: VideoModerationImageDetailListCategory{
			value: "politics",
		},
		TERRORISM: VideoModerationImageDetailListCategory{
			value: "terrorism",
		},
		PORN: VideoModerationImageDetailListCategory{
			value: "porn",
		},
		IMAGE_TEXT: VideoModerationImageDetailListCategory{
			value: "image_text",
		},
	}
}

func (c VideoModerationImageDetailListCategory) Value() string {
	return c.value
}

func (c VideoModerationImageDetailListCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoModerationImageDetailListCategory) UnmarshalJSON(b []byte) error {
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

type VideoModerationImageDetailListSuggestion struct {
	value string
}

type VideoModerationImageDetailListSuggestionEnum struct {
	BLOCK  VideoModerationImageDetailListSuggestion
	REVIEW VideoModerationImageDetailListSuggestion
}

func GetVideoModerationImageDetailListSuggestionEnum() VideoModerationImageDetailListSuggestionEnum {
	return VideoModerationImageDetailListSuggestionEnum{
		BLOCK: VideoModerationImageDetailListSuggestion{
			value: "block",
		},
		REVIEW: VideoModerationImageDetailListSuggestion{
			value: "review",
		},
	}
}

func (c VideoModerationImageDetailListSuggestion) Value() string {
	return c.value
}

func (c VideoModerationImageDetailListSuggestion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoModerationImageDetailListSuggestion) UnmarshalJSON(b []byte) error {
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
