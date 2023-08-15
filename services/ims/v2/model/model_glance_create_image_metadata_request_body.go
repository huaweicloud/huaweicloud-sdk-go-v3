package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GlanceCreateImageMetadataRequestBody 创建镜像请求体
type GlanceCreateImageMetadataRequestBody struct {

	// 镜像的操作系统具体版本,如果未指定__os_version，则默认设置为Other Linux(64 bit)，不保证该镜像能成功创建虚拟机以及通过该镜像创建的虚拟机能够正常使用。
	OsVersion *string `json:"__os_version,omitempty"`

	// 容器格式。默认取值为bare。
	ContainerFormat *string `json:"container_format,omitempty"`

	// 镜像文件格式。目前支持vhd，zvhd、zvhd2、raw，qcow2。默认取值为vhd
	DiskFormat *GlanceCreateImageMetadataRequestBodyDiskFormat `json:"disk_format,omitempty"`

	// 镜像运行需要的最小磁盘，单位为GB 。必须大于镜像系统盘容量，否则创建云主机云服务器可能失败。
	MinDisk *int32 `json:"min_disk,omitempty"`

	// 镜像运行需要的最小内存，单位为MB。参数取值依据云主机云服务器的规格限制。默认取值为0。
	MinRam *int32 `json:"min_ram,omitempty"`

	// 镜像名称，如果未指定name的取值，则默认为空，但是使用该镜像创建虚拟机会失败。名称的长度为1-255位。
	Name *string `json:"name,omitempty"`

	// 镜像是否被保护，保护后的镜像不可删除。默认取值为false。
	Protected *bool `json:"protected,omitempty"`

	// 镜像标签列表。长度为1-255位。默认为空。
	Tags *[]string `json:"tags,omitempty"`

	// 其他租户是否可见。默认取值为private。创建镜像元数据时，visibility取值只能为private。
	Visibility *string `json:"visibility,omitempty"`
}

func (o GlanceCreateImageMetadataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateImageMetadataRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceCreateImageMetadataRequestBody", string(data)}, " ")
}

type GlanceCreateImageMetadataRequestBodyDiskFormat struct {
	value string
}

type GlanceCreateImageMetadataRequestBodyDiskFormatEnum struct {
	VHD   GlanceCreateImageMetadataRequestBodyDiskFormat
	ZVHD  GlanceCreateImageMetadataRequestBodyDiskFormat
	ZVHD2 GlanceCreateImageMetadataRequestBodyDiskFormat
	RAW   GlanceCreateImageMetadataRequestBodyDiskFormat
	QCOW2 GlanceCreateImageMetadataRequestBodyDiskFormat
}

func GetGlanceCreateImageMetadataRequestBodyDiskFormatEnum() GlanceCreateImageMetadataRequestBodyDiskFormatEnum {
	return GlanceCreateImageMetadataRequestBodyDiskFormatEnum{
		VHD: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "zvhd",
		},
		ZVHD2: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "zvhd2",
		},
		RAW: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "qcow2",
		},
	}
}

func (c GlanceCreateImageMetadataRequestBodyDiskFormat) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataRequestBodyDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataRequestBodyDiskFormat) UnmarshalJSON(b []byte) error {
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
