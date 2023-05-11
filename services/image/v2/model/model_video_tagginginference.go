package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoTagginginference struct {

	// 视频标题
	VideoTitle *string `json:"video_title,omitempty"`

	// 标签语种
	Language *VideoTagginginferenceLanguage `json:"language,omitempty"`

	// 名人识别使用开关
	UseCelebrity *string `json:"use_celebrity,omitempty"`

	// 地标识别使用开关
	UseLandmark *string `json:"use_landmark,omitempty"`

	// LOGO识别使用开关
	UseLogo *string `json:"use_logo,omitempty"`

	// OCR识别使用开关
	UseOcr *string `json:"use_ocr,omitempty"`

	// 视频语音识别开关
	UseSis *string `json:"use_sis,omitempty"`

	// 图像标签识别开关
	UseTagging *string `json:"use_tagging,omitempty"`
}

func (o VideoTagginginference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTagginginference struct{}"
	}

	return strings.Join([]string{"VideoTagginginference", string(data)}, " ")
}

type VideoTagginginferenceLanguage struct {
	value string
}

type VideoTagginginferenceLanguageEnum struct {
	ZH VideoTagginginferenceLanguage
	EN VideoTagginginferenceLanguage
}

func GetVideoTagginginferenceLanguageEnum() VideoTagginginferenceLanguageEnum {
	return VideoTagginginferenceLanguageEnum{
		ZH: VideoTagginginferenceLanguage{
			value: "zh",
		},
		EN: VideoTagginginferenceLanguage{
			value: "en",
		},
	}
}

func (c VideoTagginginferenceLanguage) Value() string {
	return c.value
}

func (c VideoTagginginferenceLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoTagginginferenceLanguage) UnmarshalJSON(b []byte) error {
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
