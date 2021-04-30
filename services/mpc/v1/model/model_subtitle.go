package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Subtitle struct {
	Input *ObsObjInfo `json:"input,omitempty"`
	// 多字幕文件地址。

	Inputs *[]MulInputFileInfo `json:"inputs,omitempty"`
	// 字幕类型。取值如下：  - 0，表示不输出字幕 - 1，表示外部字幕文件嵌入视频流 - 2，表示输出WebVTT格式字幕

	SubtitleType *SubtitleSubtitleType `json:"subtitle_type,omitempty"`
}

func (o Subtitle) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Subtitle struct{}"
	}

	return strings.Join([]string{"Subtitle", string(data)}, " ")
}

type SubtitleSubtitleType struct {
	value int32
}

type SubtitleSubtitleTypeEnum struct {
	E_0 SubtitleSubtitleType
	E_1 SubtitleSubtitleType
	E_2 SubtitleSubtitleType
}

func GetSubtitleSubtitleTypeEnum() SubtitleSubtitleTypeEnum {
	return SubtitleSubtitleTypeEnum{
		E_0: SubtitleSubtitleType{
			value: 0,
		}, E_1: SubtitleSubtitleType{
			value: 1,
		}, E_2: SubtitleSubtitleType{
			value: 2,
		},
	}
}

func (c SubtitleSubtitleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SubtitleSubtitleType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
