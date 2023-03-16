package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TextToImageInference struct {

	// 文生图引导词
	Prompt string `json:"prompt"`

	// 随机数种子
	Seed *int32 `json:"seed,omitempty"`

	// 生成图片分辨率，限定字符串\"512\\*768\",\"768\\*512\",\"512\\*512\"，默认\"512\\*512\"
	Resolution *TextToImageInferenceResolution `json:"resolution,omitempty"`

	// 生成图片数量，支持1-4张，默认为1
	ImageNums *int32 `json:"image_nums,omitempty"`
}

func (o TextToImageInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextToImageInference struct{}"
	}

	return strings.Join([]string{"TextToImageInference", string(data)}, " ")
}

type TextToImageInferenceResolution struct {
	value string
}

type TextToImageInferenceResolutionEnum struct {
	E_512768 TextToImageInferenceResolution
	E_768512 TextToImageInferenceResolution
	E_512512 TextToImageInferenceResolution
}

func GetTextToImageInferenceResolutionEnum() TextToImageInferenceResolutionEnum {
	return TextToImageInferenceResolutionEnum{
		E_512768: TextToImageInferenceResolution{
			value: "512*768",
		},
		E_768512: TextToImageInferenceResolution{
			value: "768*512",
		},
		E_512512: TextToImageInferenceResolution{
			value: "512*512",
		},
	}
}

func (c TextToImageInferenceResolution) Value() string {
	return c.value
}

func (c TextToImageInferenceResolution) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TextToImageInferenceResolution) UnmarshalJSON(b []byte) error {
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
