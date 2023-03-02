package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoTranslateInference struct {

	// 字幕翻译目标语言
	TargetLanguage string `json:"target_language"`

	// 是否回写
	Rewrite *VideoTranslateInferenceRewrite `json:"rewrite,omitempty"`

	RewriteConfig *VideoTranslateInferenceRewriteConfig `json:"rewrite_config,omitempty"`
}

func (o VideoTranslateInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTranslateInference struct{}"
	}

	return strings.Join([]string{"VideoTranslateInference", string(data)}, " ")
}

type VideoTranslateInferenceRewrite struct {
	value string
}

type VideoTranslateInferenceRewriteEnum struct {
	TRUE  VideoTranslateInferenceRewrite
	FALSE VideoTranslateInferenceRewrite
}

func GetVideoTranslateInferenceRewriteEnum() VideoTranslateInferenceRewriteEnum {
	return VideoTranslateInferenceRewriteEnum{
		TRUE: VideoTranslateInferenceRewrite{
			value: "true",
		},
		FALSE: VideoTranslateInferenceRewrite{
			value: "false",
		},
	}
}

func (c VideoTranslateInferenceRewrite) Value() string {
	return c.value
}

func (c VideoTranslateInferenceRewrite) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoTranslateInferenceRewrite) UnmarshalJSON(b []byte) error {
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
