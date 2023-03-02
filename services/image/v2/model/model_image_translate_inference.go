package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImageTranslateInference struct {

	// 文本翻译目标语言，支持中文（“zh”）、翻译中文（“zh-tw”）、英语（“en”）、日语（“ja”）、泰语（“th”）、阿拉伯语（“ar”）、韩语（“ko”）
	TargetLanguage string `json:"target_language"`

	// 是否回写，默认为是
	Rewrite *ImageTranslateInferenceRewrite `json:"rewrite,omitempty"`
}

func (o ImageTranslateInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTranslateInference struct{}"
	}

	return strings.Join([]string{"ImageTranslateInference", string(data)}, " ")
}

type ImageTranslateInferenceRewrite struct {
	value string
}

type ImageTranslateInferenceRewriteEnum struct {
	TRUE  ImageTranslateInferenceRewrite
	FALSE ImageTranslateInferenceRewrite
}

func GetImageTranslateInferenceRewriteEnum() ImageTranslateInferenceRewriteEnum {
	return ImageTranslateInferenceRewriteEnum{
		TRUE: ImageTranslateInferenceRewrite{
			value: "true",
		},
		FALSE: ImageTranslateInferenceRewrite{
			value: "false",
		},
	}
}

func (c ImageTranslateInferenceRewrite) Value() string {
	return c.value
}

func (c ImageTranslateInferenceRewrite) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageTranslateInferenceRewrite) UnmarshalJSON(b []byte) error {
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
