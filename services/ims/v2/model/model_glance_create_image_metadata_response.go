package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GlanceCreateImageMetadataResponse Response Object
type GlanceCreateImageMetadataResponse struct {

	// 其他租户是否可见。取值为private。
	Visibility *string `json:"visibility,omitempty"`

	// 镜像名称，如果未指定name的取值，则默认为空，但是使用该镜像创建虚拟机会失败。名称的长度为1～128位。
	Name *string `json:"name,omitempty"`

	// 镜像是否被保护，保护后的镜像不可删除。取值为false
	Protected *bool `json:"protected,omitempty"`

	// 容器格式。取值为bare。
	ContainerFormat *string `json:"container_format,omitempty"`

	// 镜像文件格式。目前支持vhd、zvhd、raw、qcow2。默认值是vhd。
	DiskFormat *GlanceCreateImageMetadataResponseDiskFormat `json:"disk_format,omitempty"`

	// 镜像标签列表。长度为1～255位。
	Tags *[]string `json:"tags,omitempty"`

	// 镜像运行最小内存，单位为MB。取值参考ECS规格限制，一般设置为0。云服务器的规格限制，请参见规格清单。
	MinRam *int32 `json:"min_ram,omitempty"`

	// 镜像运行需要的最小磁盘容量，单位为GB 。取值为40～1024GB。必须大于镜像系统盘容量，否则创建云主机云服务器可能失败。
	MinDisk *int32 `json:"min_disk,omitempty"`

	// 镜像状态。取值如下：queued：表示镜像元数据已经创建成功，等待上传镜像文件。saving：表示镜像正在上传文件到后端存储。deleted：表示镜像已经删除。killed：表示镜像上传错误。active：表示镜像可以正常使用。
	Status *GlanceCreateImageMetadataResponseStatus `json:"status,omitempty"`

	// 创建时间。格式为UTC时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。格式为UTC时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 本镜像链接。
	Self *string `json:"self,omitempty"`

	// 镜像ID，用户调用创建镜像接口后，需保存该镜像的ID，用来调用上传镜像接口完成镜像上传。
	Id *string `json:"id,omitempty"`

	// 上传下载镜像文件的地址链接。
	File *string `json:"file,omitempty"`

	// 视图链接。
	Schema *string `json:"schema,omitempty"`

	// 镜像后端存储类型，目前支持uds。
	ImageSourceType *string `json:"__image_source_type,omitempty"`

	// 镜像大小。单位为字节。
	ImageSize *string `json:"__image_size,omitempty"`

	// 镜像是否注册。只有已注册的镜像才能在Portal界面上查询到。取值为true。
	Isregistered *string `json:"__isregistered,omitempty"`

	// 镜像的操作系统具体版本。
	OsVersion *string `json:"__os_version,omitempty"`

	// 镜像的操作系统类型，取值由__os_version确定。支持Windows、Linux和other。
	OsType *GlanceCreateImageMetadataResponseOsType `json:"__os_type,omitempty"`

	// 表示镜像支持的操作系统平台。取值由__os_version确定
	Platform *string `json:"__platform,omitempty"`

	// 表示操作系统位数。取值由__os_version确定，取值为32或64。
	OsBit *GlanceCreateImageMetadataResponseOsBit `json:"__os_bit,omitempty"`

	// 镜像类型。取值为private，表示私有镜像。
	Imagetype *string `json:"__imagetype,omitempty"`

	// 平台类型。镜像使用环境类型：FusionCompute、Ironic、DataImage。如果是云主机云服务器镜像，则取值为FusionCompute。如果是数据卷镜像则取值是DataImage。如果是物理机裸金属服务器镜像，则取值是Ironic。
	VirtualEnvType *GlanceCreateImageMetadataResponseVirtualEnvType `json:"virtual_env_type,omitempty"`

	// 镜像所属项目ID。
	Owner *string `json:"owner,omitempty"`

	// 镜像虚拟大小。单位为字节。
	VirtualSize *int32 `json:"virtual_size,omitempty"`

	// 镜像属性的集合，不表示具体的镜像属性
	Properties *interface{} `json:"properties,omitempty"`

	// 表示当前镜像来源是从外部导入。取值：file
	RootOrigin *string `json:"__root_origin,omitempty"`

	// 镜像文件md5值。
	Checksum *string `json:"checksum,omitempty"`

	// 目前暂时不使用。
	Size           *int64 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o GlanceCreateImageMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateImageMetadataResponse struct{}"
	}

	return strings.Join([]string{"GlanceCreateImageMetadataResponse", string(data)}, " ")
}

type GlanceCreateImageMetadataResponseDiskFormat struct {
	value string
}

type GlanceCreateImageMetadataResponseDiskFormatEnum struct {
	VHD   GlanceCreateImageMetadataResponseDiskFormat
	ZVHD  GlanceCreateImageMetadataResponseDiskFormat
	RAW   GlanceCreateImageMetadataResponseDiskFormat
	QCOW2 GlanceCreateImageMetadataResponseDiskFormat
}

func GetGlanceCreateImageMetadataResponseDiskFormatEnum() GlanceCreateImageMetadataResponseDiskFormatEnum {
	return GlanceCreateImageMetadataResponseDiskFormatEnum{
		VHD: GlanceCreateImageMetadataResponseDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceCreateImageMetadataResponseDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceCreateImageMetadataResponseDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceCreateImageMetadataResponseDiskFormat{
			value: "qcow2",
		},
	}
}

func (c GlanceCreateImageMetadataResponseDiskFormat) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceCreateImageMetadataResponseStatus struct {
	value string
}

type GlanceCreateImageMetadataResponseStatusEnum struct {
	QUEUED  GlanceCreateImageMetadataResponseStatus
	SAVING  GlanceCreateImageMetadataResponseStatus
	DELETED GlanceCreateImageMetadataResponseStatus
	KILLED  GlanceCreateImageMetadataResponseStatus
	ACTIVE  GlanceCreateImageMetadataResponseStatus
}

func GetGlanceCreateImageMetadataResponseStatusEnum() GlanceCreateImageMetadataResponseStatusEnum {
	return GlanceCreateImageMetadataResponseStatusEnum{
		QUEUED: GlanceCreateImageMetadataResponseStatus{
			value: "queued",
		},
		SAVING: GlanceCreateImageMetadataResponseStatus{
			value: "saving",
		},
		DELETED: GlanceCreateImageMetadataResponseStatus{
			value: "deleted",
		},
		KILLED: GlanceCreateImageMetadataResponseStatus{
			value: "killed",
		},
		ACTIVE: GlanceCreateImageMetadataResponseStatus{
			value: "active",
		},
	}
}

func (c GlanceCreateImageMetadataResponseStatus) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseStatus) UnmarshalJSON(b []byte) error {
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

type GlanceCreateImageMetadataResponseOsType struct {
	value string
}

type GlanceCreateImageMetadataResponseOsTypeEnum struct {
	WINDOWS GlanceCreateImageMetadataResponseOsType
	LINUX   GlanceCreateImageMetadataResponseOsType
	OTHER   GlanceCreateImageMetadataResponseOsType
}

func GetGlanceCreateImageMetadataResponseOsTypeEnum() GlanceCreateImageMetadataResponseOsTypeEnum {
	return GlanceCreateImageMetadataResponseOsTypeEnum{
		WINDOWS: GlanceCreateImageMetadataResponseOsType{
			value: "Windows",
		},
		LINUX: GlanceCreateImageMetadataResponseOsType{
			value: "Linux",
		},
		OTHER: GlanceCreateImageMetadataResponseOsType{
			value: "other",
		},
	}
}

func (c GlanceCreateImageMetadataResponseOsType) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseOsType) UnmarshalJSON(b []byte) error {
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

type GlanceCreateImageMetadataResponseOsBit struct {
	value string
}

type GlanceCreateImageMetadataResponseOsBitEnum struct {
	E_32 GlanceCreateImageMetadataResponseOsBit
	E_64 GlanceCreateImageMetadataResponseOsBit
}

func GetGlanceCreateImageMetadataResponseOsBitEnum() GlanceCreateImageMetadataResponseOsBitEnum {
	return GlanceCreateImageMetadataResponseOsBitEnum{
		E_32: GlanceCreateImageMetadataResponseOsBit{
			value: "32",
		},
		E_64: GlanceCreateImageMetadataResponseOsBit{
			value: "64",
		},
	}
}

func (c GlanceCreateImageMetadataResponseOsBit) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceCreateImageMetadataResponseVirtualEnvType struct {
	value string
}

type GlanceCreateImageMetadataResponseVirtualEnvTypeEnum struct {
	FUSION_COMPUTE GlanceCreateImageMetadataResponseVirtualEnvType
	IRONIC         GlanceCreateImageMetadataResponseVirtualEnvType
	DATA_IMAGE     GlanceCreateImageMetadataResponseVirtualEnvType
}

func GetGlanceCreateImageMetadataResponseVirtualEnvTypeEnum() GlanceCreateImageMetadataResponseVirtualEnvTypeEnum {
	return GlanceCreateImageMetadataResponseVirtualEnvTypeEnum{
		FUSION_COMPUTE: GlanceCreateImageMetadataResponseVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: GlanceCreateImageMetadataResponseVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: GlanceCreateImageMetadataResponseVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c GlanceCreateImageMetadataResponseVirtualEnvType) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseVirtualEnvType) UnmarshalJSON(b []byte) error {
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
