package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImageHighresolutionMattingInference struct {

	// 是否只返回处理结果的alpha通道，\"foreground\"代表返回带alpha通道的前景图片，\"alpha\"字符串代表仅返回alpha通道
	ReturnType ImageHighresolutionMattingInferenceReturnType `json:"return_type"`

	// 指定抠图区域坐标，默认全图，示例：[x_min,y_min,x_max,y_max]
	Coord *[]int32 `json:"coord,omitempty"`
}

func (o ImageHighresolutionMattingInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageHighresolutionMattingInference struct{}"
	}

	return strings.Join([]string{"ImageHighresolutionMattingInference", string(data)}, " ")
}

type ImageHighresolutionMattingInferenceReturnType struct {
	value string
}

type ImageHighresolutionMattingInferenceReturnTypeEnum struct {
	FOREGROUND ImageHighresolutionMattingInferenceReturnType
	ALPHA      ImageHighresolutionMattingInferenceReturnType
}

func GetImageHighresolutionMattingInferenceReturnTypeEnum() ImageHighresolutionMattingInferenceReturnTypeEnum {
	return ImageHighresolutionMattingInferenceReturnTypeEnum{
		FOREGROUND: ImageHighresolutionMattingInferenceReturnType{
			value: "foreground",
		},
		ALPHA: ImageHighresolutionMattingInferenceReturnType{
			value: "alpha",
		},
	}
}

func (c ImageHighresolutionMattingInferenceReturnType) Value() string {
	return c.value
}

func (c ImageHighresolutionMattingInferenceReturnType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageHighresolutionMattingInferenceReturnType) UnmarshalJSON(b []byte) error {
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
