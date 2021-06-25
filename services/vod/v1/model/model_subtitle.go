package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Subtitle struct {
	// 字幕文件id<br/>

	Id int32 `json:"id"`
	// 字幕文件类型<br/>

	Type SubtitleType `json:"type"`
	// 字幕文件语音种类<br/>

	Language SubtitleLanguage `json:"language"`
	// 字幕文件md5值<br/>

	Md5 *string `json:"md5,omitempty"`
	// 字幕文件描述<br/>

	Description *string `json:"description,omitempty"`
}

func (o Subtitle) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Subtitle struct{}"
	}

	return strings.Join([]string{"Subtitle", string(data)}, " ")
}

type SubtitleType struct {
	value string
}

type SubtitleTypeEnum struct {
	SRT SubtitleType
}

func GetSubtitleTypeEnum() SubtitleTypeEnum {
	return SubtitleTypeEnum{
		SRT: SubtitleType{
			value: "SRT",
		},
	}
}

func (c SubtitleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SubtitleType) UnmarshalJSON(b []byte) error {
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

type SubtitleLanguage struct {
	value string
}

type SubtitleLanguageEnum struct {
	CN SubtitleLanguage
	EN SubtitleLanguage
}

func GetSubtitleLanguageEnum() SubtitleLanguageEnum {
	return SubtitleLanguageEnum{
		CN: SubtitleLanguage{
			value: "CN",
		},
		EN: SubtitleLanguage{
			value: "EN",
		},
	}
}

func (c SubtitleLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SubtitleLanguage) UnmarshalJSON(b []byte) error {
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
