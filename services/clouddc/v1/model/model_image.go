package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Image 服务器镜像信息。
type Image struct {

	// 镜像ID，格式为UUID。
	Id string `json:"id"`

	// 镜像名称
	Name string `json:"name"`

	// 镜像os类型
	OsType ImageOsType `json:"os_type"`
}

func (o Image) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Image struct{}"
	}

	return strings.Join([]string{"Image", string(data)}, " ")
}

type ImageOsType struct {
	value string
}

type ImageOsTypeEnum struct {
	LINUX   ImageOsType
	WINDOWS ImageOsType
	OTHER   ImageOsType
}

func GetImageOsTypeEnum() ImageOsTypeEnum {
	return ImageOsTypeEnum{
		LINUX: ImageOsType{
			value: "Linux",
		},
		WINDOWS: ImageOsType{
			value: "Windows",
		},
		OTHER: ImageOsType{
			value: "Other",
		},
	}
}

func (c ImageOsType) Value() string {
	return c.value
}

func (c ImageOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageOsType) UnmarshalJSON(b []byte) error {
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
