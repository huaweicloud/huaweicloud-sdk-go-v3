package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImageType struct {
	value string
}

type ImageTypeEnum struct {
	APP      ImageType
	NOTEBOOK ImageType
}

func GetImageTypeEnum() ImageTypeEnum {
	return ImageTypeEnum{
		APP: ImageType{
			value: "APP",
		},
		NOTEBOOK: ImageType{
			value: "NOTEBOOK",
		},
	}
}

func (c ImageType) Value() string {
	return c.value
}

func (c ImageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageType) UnmarshalJSON(b []byte) error {
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
