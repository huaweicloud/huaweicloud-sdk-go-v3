package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageTypeEnum 镜像类型： * `gold` - 云市场镜像 * `public` - 公共镜像 * `private` - 私有镜像 * `shared` - 共享镜像 * `other` - 其他
type ImageTypeEnum struct {
	value string
}

type ImageTypeEnumEnum struct {
	GOLD    ImageTypeEnum
	PUBLIC  ImageTypeEnum
	PRIVATE ImageTypeEnum
	SHARED  ImageTypeEnum
	OTHER   ImageTypeEnum
}

func GetImageTypeEnumEnum() ImageTypeEnumEnum {
	return ImageTypeEnumEnum{
		GOLD: ImageTypeEnum{
			value: "gold",
		},
		PUBLIC: ImageTypeEnum{
			value: "public",
		},
		PRIVATE: ImageTypeEnum{
			value: "private",
		},
		SHARED: ImageTypeEnum{
			value: "shared",
		},
		OTHER: ImageTypeEnum{
			value: "other",
		},
	}
}

func (c ImageTypeEnum) Value() string {
	return c.value
}

func (c ImageTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageTypeEnum) UnmarshalJSON(b []byte) error {
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
