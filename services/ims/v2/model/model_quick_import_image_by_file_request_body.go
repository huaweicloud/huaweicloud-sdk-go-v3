package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QuickImportImageByFileRequestBody 快速通道创建镜像的请求体
type QuickImportImageByFileRequestBody struct {

	// 镜像名称
	Name string `json:"name"`

	// 镜像描述信息。_description参数说明请参考镜像属性。支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。默认为空。
	Description *string `json:"description,omitempty"`

	// 操作系统版本。使用上传至OBS桶中的外部镜像文件制作镜像时生效
	OsVersion string `json:"os_version"`

	// OBS桶中外部镜像文件地址。在使用OBS桶的外部镜像文件制作镜像时生效且为必选字段。格式为<OBS桶名>:<OBS镜像文件名称>。注意：此处的OBS桶和镜像文件的存储类别必须是OBS标准存储。
	ImageUrl string `json:"image_url"`

	// 最小系统盘大小。在使用OBS桶的外部镜像文件制作镜像时生效且为必选字段。取值为1至1024GB。
	MinDisk int32 `json:"min_disk"`

	// 镜像标签列表。默认为空。 tags和image_tags只能使用一个。
	Tags *[]string `json:"tags,omitempty"`

	// 制作的镜像类型。系统盘镜像为ECS/BMS，数据盘镜像为DataImage. 制作数据盘镜像时该参数必选.
	Type *QuickImportImageByFileRequestBodyType `json:"type,omitempty"`

	// 表示当前镜像所属的企业项目。 取值为0或无该值，表示属于default企业项目。 取值为UUID，表示属于该UUID对应的企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像的架构类型。取值包括： x86 arm 默认使用“x86”。
	Architecture *QuickImportImageByFileRequestBodyArchitecture `json:"architecture,omitempty"`

	// 操作系统版本。 创建数据盘镜像时该参数取值为Linux或Windows，默认Linux。
	OsType *QuickImportImageByFileRequestBodyOsType `json:"os_type,omitempty"`

	// 新规范的镜像标签列表。默认为空。 tags和image_tags只能使用一个。
	ImageTags *[]ResourceTag `json:"image_tags,omitempty"`
}

func (o QuickImportImageByFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuickImportImageByFileRequestBody struct{}"
	}

	return strings.Join([]string{"QuickImportImageByFileRequestBody", string(data)}, " ")
}

type QuickImportImageByFileRequestBodyType struct {
	value string
}

type QuickImportImageByFileRequestBodyTypeEnum struct {
	ECS        QuickImportImageByFileRequestBodyType
	BMS        QuickImportImageByFileRequestBodyType
	DATA_IMAGE QuickImportImageByFileRequestBodyType
}

func GetQuickImportImageByFileRequestBodyTypeEnum() QuickImportImageByFileRequestBodyTypeEnum {
	return QuickImportImageByFileRequestBodyTypeEnum{
		ECS: QuickImportImageByFileRequestBodyType{
			value: "ECS",
		},
		BMS: QuickImportImageByFileRequestBodyType{
			value: "BMS",
		},
		DATA_IMAGE: QuickImportImageByFileRequestBodyType{
			value: "DataImage",
		},
	}
}

func (c QuickImportImageByFileRequestBodyType) Value() string {
	return c.value
}

func (c QuickImportImageByFileRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuickImportImageByFileRequestBodyType) UnmarshalJSON(b []byte) error {
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

type QuickImportImageByFileRequestBodyArchitecture struct {
	value string
}

type QuickImportImageByFileRequestBodyArchitectureEnum struct {
	X86 QuickImportImageByFileRequestBodyArchitecture
	ARM QuickImportImageByFileRequestBodyArchitecture
}

func GetQuickImportImageByFileRequestBodyArchitectureEnum() QuickImportImageByFileRequestBodyArchitectureEnum {
	return QuickImportImageByFileRequestBodyArchitectureEnum{
		X86: QuickImportImageByFileRequestBodyArchitecture{
			value: "x86",
		},
		ARM: QuickImportImageByFileRequestBodyArchitecture{
			value: "arm",
		},
	}
}

func (c QuickImportImageByFileRequestBodyArchitecture) Value() string {
	return c.value
}

func (c QuickImportImageByFileRequestBodyArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuickImportImageByFileRequestBodyArchitecture) UnmarshalJSON(b []byte) error {
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

type QuickImportImageByFileRequestBodyOsType struct {
	value string
}

type QuickImportImageByFileRequestBodyOsTypeEnum struct {
	LINUX   QuickImportImageByFileRequestBodyOsType
	WINDOWS QuickImportImageByFileRequestBodyOsType
}

func GetQuickImportImageByFileRequestBodyOsTypeEnum() QuickImportImageByFileRequestBodyOsTypeEnum {
	return QuickImportImageByFileRequestBodyOsTypeEnum{
		LINUX: QuickImportImageByFileRequestBodyOsType{
			value: "Linux",
		},
		WINDOWS: QuickImportImageByFileRequestBodyOsType{
			value: "Windows",
		},
	}
}

func (c QuickImportImageByFileRequestBodyOsType) Value() string {
	return c.value
}

func (c QuickImportImageByFileRequestBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuickImportImageByFileRequestBodyOsType) UnmarshalJSON(b []byte) error {
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
