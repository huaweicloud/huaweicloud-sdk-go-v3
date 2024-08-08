package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageRef 镜像源信息。
type ImageRef struct {

	// 镜像源的唯一标识。
	Id string `json:"id"`

	// 镜像源的镜像类型: * `gold` - 云市场镜像 * `public` - 公共镜像 * `private` - 私有镜像 * `shared` - 共享镜像 * `other` - 其他
	ImageType ImageRefImageType `json:"image_type"`

	// 镜像源的规格编码，对于`gold`镜像类型，这个值是的必须项。
	SpceCode *string `json:"spce_code,omitempty"`

	// 镜像源的产品ID，对于`gold`镜像类型，这个值是的必须项。
	ProductId *string `json:"product_id,omitempty"`
}

func (o ImageRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageRef struct{}"
	}

	return strings.Join([]string{"ImageRef", string(data)}, " ")
}

type ImageRefImageType struct {
	value string
}

type ImageRefImageTypeEnum struct {
	GOLD    ImageRefImageType
	PUBLIC  ImageRefImageType
	PRIVATE ImageRefImageType
	SHARED  ImageRefImageType
	OTHER   ImageRefImageType
}

func GetImageRefImageTypeEnum() ImageRefImageTypeEnum {
	return ImageRefImageTypeEnum{
		GOLD: ImageRefImageType{
			value: "gold",
		},
		PUBLIC: ImageRefImageType{
			value: "public",
		},
		PRIVATE: ImageRefImageType{
			value: "private",
		},
		SHARED: ImageRefImageType{
			value: "shared",
		},
		OTHER: ImageRefImageType{
			value: "other",
		},
	}
}

func (c ImageRefImageType) Value() string {
	return c.value
}

func (c ImageRefImageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageRefImageType) UnmarshalJSON(b []byte) error {
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
