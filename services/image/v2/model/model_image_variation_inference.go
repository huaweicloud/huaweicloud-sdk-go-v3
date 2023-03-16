package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImageVariationInference struct {

	// 随机数种子
	Seed *int32 `json:"seed,omitempty"`

	// 生成图片分辨率，限定字符串\"512\\*768\",\"768\\*512\",\"512\\*512\"，默认\"512\\*512\"
	Resolution *ImageVariationInferenceResolution `json:"resolution,omitempty"`

	// 生成图片数量，支持1-4张，默认为1
	ImageNums *int32 `json:"image_nums,omitempty"`
}

func (o ImageVariationInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageVariationInference struct{}"
	}

	return strings.Join([]string{"ImageVariationInference", string(data)}, " ")
}

type ImageVariationInferenceResolution struct {
	value string
}

type ImageVariationInferenceResolutionEnum struct {
	E_512768 ImageVariationInferenceResolution
	E_768512 ImageVariationInferenceResolution
	E_512512 ImageVariationInferenceResolution
}

func GetImageVariationInferenceResolutionEnum() ImageVariationInferenceResolutionEnum {
	return ImageVariationInferenceResolutionEnum{
		E_512768: ImageVariationInferenceResolution{
			value: "512*768",
		},
		E_768512: ImageVariationInferenceResolution{
			value: "768*512",
		},
		E_512512: ImageVariationInferenceResolution{
			value: "512*512",
		},
	}
}

func (c ImageVariationInferenceResolution) Value() string {
	return c.value
}

func (c ImageVariationInferenceResolution) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageVariationInferenceResolution) UnmarshalJSON(b []byte) error {
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
