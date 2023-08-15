package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImageInfo struct {

	// 镜像ID。
	Id *string `json:"id,omitempty"`

	// 镜像类型，目前支持以下类型： 公共镜像：gold 私有镜像：private。
	ImageType *ImageInfoImageType `json:"image_type,omitempty"`

	// 操作系统类型，目前取值Linux， Windows，Other。
	OsType *string `json:"os_type,omitempty"`

	// 操作系统具体版本。
	OsVersion *string `json:"os_version,omitempty"`

	// 镜像格式，目前支持vhd，raw，qcow2，zvhd2格式。
	DiskFormat *string `json:"disk_format,omitempty"`

	// 镜像名称。
	Name *string `json:"name,omitempty"`

	// 镜像运行需要的最小内存，单位为MB。参数取值依据弹性云服务器的规格限制，一般设置为0。
	MinRam *int32 `json:"min_ram,omitempty"`

	// 镜像运行需要的最小磁盘，单位为GB 。取值为40～1024GB。
	MinDisk *int32 `json:"min_disk,omitempty"`

	// 镜像的产品编码。
	ProductCode *string `json:"product_code,omitempty"`
}

func (o ImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageInfo struct{}"
	}

	return strings.Join([]string{"ImageInfo", string(data)}, " ")
}

type ImageInfoImageType struct {
	value string
}

type ImageInfoImageTypeEnum struct {
	GOLD    ImageInfoImageType
	PRIVATE ImageInfoImageType
}

func GetImageInfoImageTypeEnum() ImageInfoImageTypeEnum {
	return ImageInfoImageTypeEnum{
		GOLD: ImageInfoImageType{
			value: "gold",
		},
		PRIVATE: ImageInfoImageType{
			value: "private",
		},
	}
}

func (c ImageInfoImageType) Value() string {
	return c.value
}

func (c ImageInfoImageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoImageType) UnmarshalJSON(b []byte) error {
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
