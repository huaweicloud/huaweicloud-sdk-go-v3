package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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
