package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建镜像请求参数体
type CreateImageRequestBody struct {
	// 需要转换的数据盘信息，其中，当使用云服务器上的数据盘进行私有数据盘镜像创建时，该字段必选。 如果不是用于制作数据盘镜像，该字段默认为空。

	DataImages *[]CreateDataImage `json:"data_images,omitempty"`
	// 镜像描述信息。支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。默认为空。

	Description *string `json:"description,omitempty"`
	// 表示当前镜像所属的企业项目。取值为0或无该值，表示属于default企业项目。取值为UUID，表示属于该UUID对应的企业项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 新规范的镜像标签列表。默认为空。tags和image_tags只能使用一个。

	ImageTags *[]TagKeyValue `json:"image_tags,omitempty"`
	// 需要转换的云服务器ID。使用instance_id字段，从云服务器制作私有镜像时，该字段填写云服务器ID。

	InstanceId *string `json:"instance_id,omitempty"`
	// 镜像名称

	Name string `json:"name"`
	// 镜像标签列表。默认为空。tags和image_tags只能使用一个。

	Tags *[]string `json:"tags,omitempty"`
	// 表示镜像支持的最大内存，单位为MB。

	MaxRam *int32 `json:"max_ram,omitempty"`
	// 表示镜像支持的最小内存，单位为MB，默认为0，表示不受限制。

	MinRam *int32 `json:"min_ram,omitempty"`
	// 操作系统版本。 使用上传至OBS桶中的外部镜像文件制作镜像时生效。 当“is_quick_import”的值为“true”时，即使用镜像文件快速导入方式导入系统盘镜像，则该参数为必填参数。

	OsVersion *string `json:"os_version,omitempty"`
	// OBS桶中外部镜像文件地址。 在使用OBS桶的外部镜像文件制作镜像时生效且为必选字段。格式为<OBS桶名>:<OBS镜像文件名称>。

	ImageUrl *string `json:"image_url,omitempty"`
	// 最小系统盘大小。 在使用OBS桶的外部镜像文件制作镜像时生效且为必选字段。取值为40～1024GB。

	MinDisk *int32 `json:"min_disk,omitempty"`
	// 是否自动配置。 取值为true或false。 如果需要后台自动配置，取值为true，否则为false。默认取值为false。

	IsConfig *bool `json:"is_config,omitempty"`
	// 创建加密镜像的用户主密钥，具体取值请参考《密钥管理服务用户指南》获取。

	CmkId *string `json:"cmk_id,omitempty"`
	// 镜像的类型。 取值为ECS、BMS、FusionCompute、Ironic。默认使用“ECS”。 ECS/FusionCompute：表示是ECS服务器的镜像。 BMS/Ironic：表示是BMS服务器的镜像。

	Type *CreateImageRequestBodyType `json:"type,omitempty"`
	// 是否使用镜像文件快速导入方式，导入系统盘镜像。 是，配置为true。 否，配置为false。 关于镜像文件快速导入的约束与限制请参见镜像文件快速导入。

	IsQuickImport *bool `json:"is_quick_import,omitempty"`
	// 镜像的架构类型。取值包括： x86 arm 默认使用“x86”。 当架构类型为arm时，镜像引导方式将自动转为UEFI的引导方式。

	Architecture *CreateImageRequestBodyArchitecture `json:"architecture,omitempty"`
	// 数据盘的卷ID。当数据盘创建系统盘镜像时，该参数必选

	VolumeId *string `json:"volume_id,omitempty"`
}

func (o CreateImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageRequestBody", string(data)}, " ")
}

type CreateImageRequestBodyType struct {
	value string
}

type CreateImageRequestBodyTypeEnum struct {
	ECS            CreateImageRequestBodyType
	BMS            CreateImageRequestBodyType
	FUSION_COMPUTE CreateImageRequestBodyType
	IRONIC         CreateImageRequestBodyType
}

func GetCreateImageRequestBodyTypeEnum() CreateImageRequestBodyTypeEnum {
	return CreateImageRequestBodyTypeEnum{
		ECS: CreateImageRequestBodyType{
			value: "ECS",
		},
		BMS: CreateImageRequestBodyType{
			value: "BMS",
		},
		FUSION_COMPUTE: CreateImageRequestBodyType{
			value: "FusionCompute",
		},
		IRONIC: CreateImageRequestBodyType{
			value: "Ironic",
		},
	}
}

func (c CreateImageRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageRequestBodyType) UnmarshalJSON(b []byte) error {
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

type CreateImageRequestBodyArchitecture struct {
	value string
}

type CreateImageRequestBodyArchitectureEnum struct {
	X86 CreateImageRequestBodyArchitecture
	ARM CreateImageRequestBodyArchitecture
}

func GetCreateImageRequestBodyArchitectureEnum() CreateImageRequestBodyArchitectureEnum {
	return CreateImageRequestBodyArchitectureEnum{
		X86: CreateImageRequestBodyArchitecture{
			value: "x86",
		},
		ARM: CreateImageRequestBodyArchitecture{
			value: "arm",
		},
	}
}

func (c CreateImageRequestBodyArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageRequestBodyArchitecture) UnmarshalJSON(b []byte) error {
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
