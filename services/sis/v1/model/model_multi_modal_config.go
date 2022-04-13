package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 多模态评测的配置
type MultiModalConfig struct {
	// 视频的封装格式。不填写此字段，则默认为auto。注意不论何种格式，均要求帧率在25fps以上，清晰度在240*240以上。   auto  自动判断，系统会自动判断视频封装格式。  avi  avi封装格式。  mp4  mp4封装格式。  webm  webm封装格式。  mkv  mkv封装格式。  flv  flv封装格式。

	VideoFormat *MultiModalConfigVideoFormat `json:"video_format,omitempty"`
	// 评测语言和口音。  en_gb 英语-英式口音。

	Language MultiModalConfigLanguage `json:"language"`
	// 评测模式。  word 单词模式。  sentence 句子模式。

	Mode MultiModalConfigMode `json:"mode"`
}

func (o MultiModalConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiModalConfig struct{}"
	}

	return strings.Join([]string{"MultiModalConfig", string(data)}, " ")
}

type MultiModalConfigVideoFormat struct {
	value string
}

type MultiModalConfigVideoFormatEnum struct {
	AUTO MultiModalConfigVideoFormat
	AVI  MultiModalConfigVideoFormat
	MP4  MultiModalConfigVideoFormat
	WEBM MultiModalConfigVideoFormat
	MKV  MultiModalConfigVideoFormat
	FLV  MultiModalConfigVideoFormat
}

func GetMultiModalConfigVideoFormatEnum() MultiModalConfigVideoFormatEnum {
	return MultiModalConfigVideoFormatEnum{
		AUTO: MultiModalConfigVideoFormat{
			value: "auto",
		},
		AVI: MultiModalConfigVideoFormat{
			value: "avi",
		},
		MP4: MultiModalConfigVideoFormat{
			value: "mp4",
		},
		WEBM: MultiModalConfigVideoFormat{
			value: "webm",
		},
		MKV: MultiModalConfigVideoFormat{
			value: "mkv",
		},
		FLV: MultiModalConfigVideoFormat{
			value: "flv",
		},
	}
}

func (c MultiModalConfigVideoFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiModalConfigVideoFormat) UnmarshalJSON(b []byte) error {
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

type MultiModalConfigLanguage struct {
	value string
}

type MultiModalConfigLanguageEnum struct {
	EN_GB MultiModalConfigLanguage
}

func GetMultiModalConfigLanguageEnum() MultiModalConfigLanguageEnum {
	return MultiModalConfigLanguageEnum{
		EN_GB: MultiModalConfigLanguage{
			value: "en_gb",
		},
	}
}

func (c MultiModalConfigLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiModalConfigLanguage) UnmarshalJSON(b []byte) error {
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

type MultiModalConfigMode struct {
	value string
}

type MultiModalConfigModeEnum struct {
	WORD     MultiModalConfigMode
	SENTENCE MultiModalConfigMode
}

func GetMultiModalConfigModeEnum() MultiModalConfigModeEnum {
	return MultiModalConfigModeEnum{
		WORD: MultiModalConfigMode{
			value: "word",
		},
		SENTENCE: MultiModalConfigMode{
			value: "sentence",
		},
	}
}

func (c MultiModalConfigMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiModalConfigMode) UnmarshalJSON(b []byte) error {
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
