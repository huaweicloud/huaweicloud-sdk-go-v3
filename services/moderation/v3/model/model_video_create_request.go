package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoCreateRequest struct {
	Data *VideoCreateRequestData `json:"data"`

	// 事件类型，可选值如下： default：默认事件
	EventType VideoCreateRequestEventType `json:"event_type"`

	// 视频中画面需要检测的风险类型，列表不能为空。 terrorism：涉政暴恐内容的检测 porn：鉴黄内容的检测 politics：政治敏感人物内容的检测 image_text：图文违规内容的检测（检测图片中出现的广告、色情、暴恐、涉政的文字违规内容以及二维码内容）
	ImageCategories []VideoCreateRequestImageCategories `json:"image_categories"`

	// 视频中音频需要检测的风险类型，不传或为空时表示不审核音频维度。 politics: 涉政检测 porn：涉黄检测 ad: 广告检测 moan: 娇喘检测 abuse: 辱骂检测
	AudioCategories *[]VideoCreateRequestAudioCategories `json:"audio_categories,omitempty"`

	// 回调http接口：当该字段非空时，服务将根据该字段回调通知用户审核结果。
	Callback *string `json:"callback,omitempty"`
}

func (o VideoCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCreateRequest struct{}"
	}

	return strings.Join([]string{"VideoCreateRequest", string(data)}, " ")
}

type VideoCreateRequestEventType struct {
	value string
}

type VideoCreateRequestEventTypeEnum struct {
	DEFAULT VideoCreateRequestEventType
}

func GetVideoCreateRequestEventTypeEnum() VideoCreateRequestEventTypeEnum {
	return VideoCreateRequestEventTypeEnum{
		DEFAULT: VideoCreateRequestEventType{
			value: "default",
		},
	}
}

func (c VideoCreateRequestEventType) Value() string {
	return c.value
}

func (c VideoCreateRequestEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoCreateRequestEventType) UnmarshalJSON(b []byte) error {
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

type VideoCreateRequestImageCategories struct {
	value string
}

type VideoCreateRequestImageCategoriesEnum struct {
	PORN       VideoCreateRequestImageCategories
	POLITICS   VideoCreateRequestImageCategories
	TERRORISM  VideoCreateRequestImageCategories
	IMAGE_TEXT VideoCreateRequestImageCategories
}

func GetVideoCreateRequestImageCategoriesEnum() VideoCreateRequestImageCategoriesEnum {
	return VideoCreateRequestImageCategoriesEnum{
		PORN: VideoCreateRequestImageCategories{
			value: "porn",
		},
		POLITICS: VideoCreateRequestImageCategories{
			value: "politics",
		},
		TERRORISM: VideoCreateRequestImageCategories{
			value: "terrorism",
		},
		IMAGE_TEXT: VideoCreateRequestImageCategories{
			value: "image_text",
		},
	}
}

func (c VideoCreateRequestImageCategories) Value() string {
	return c.value
}

func (c VideoCreateRequestImageCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoCreateRequestImageCategories) UnmarshalJSON(b []byte) error {
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

type VideoCreateRequestAudioCategories struct {
	value string
}

type VideoCreateRequestAudioCategoriesEnum struct {
	PORN     VideoCreateRequestAudioCategories
	AD       VideoCreateRequestAudioCategories
	POLITICS VideoCreateRequestAudioCategories
	MOAN     VideoCreateRequestAudioCategories
	ABUSE    VideoCreateRequestAudioCategories
}

func GetVideoCreateRequestAudioCategoriesEnum() VideoCreateRequestAudioCategoriesEnum {
	return VideoCreateRequestAudioCategoriesEnum{
		PORN: VideoCreateRequestAudioCategories{
			value: "porn",
		},
		AD: VideoCreateRequestAudioCategories{
			value: "ad",
		},
		POLITICS: VideoCreateRequestAudioCategories{
			value: "politics",
		},
		MOAN: VideoCreateRequestAudioCategories{
			value: "moan",
		},
		ABUSE: VideoCreateRequestAudioCategories{
			value: "abuse",
		},
	}
}

func (c VideoCreateRequestAudioCategories) Value() string {
	return c.value
}

func (c VideoCreateRequestAudioCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoCreateRequestAudioCategories) UnmarshalJSON(b []byte) error {
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
