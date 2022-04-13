package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 镜像信息响应体
type ImageInfo struct {
	// 备份ID。如果是备份创建的镜像，则填写为备份的ID，否则为空

	BackupId *string `json:"__backup_id,omitempty"`
	// 镜像来源。公共镜像为空

	DataOrigin *string `json:"__data_origin,omitempty"`
	// 镜像描述信息。 支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。

	Description *string `json:"__description,omitempty"`
	// 镜像文件的大小，单位为字节

	ImageSize string `json:"__image_size"`
	// 镜像后端存储类型，目前只支持uds

	ImageSourceType ImageInfoImageSourceType `json:"__image_source_type"`
	// 镜像类型，目前支持以下类型： 公共镜像：gold 私有镜像：private 共享镜像：shared

	Imagetype ImageInfoImagetype `json:"__imagetype"`
	// 是否是注册过的镜像，取值为“true”或者“false”

	Isregistered ImageInfoIsregistered `json:"__isregistered"`
	// 父镜像ID。公共镜像或通过文件创建的私有镜像，取值为空

	Originalimagename *string `json:"__originalimagename,omitempty"`
	// 操作系统位数，一般取值为“32”或者“64”

	OsBit *ImageInfoOsBit `json:"__os_bit,omitempty"`
	// 操作系统类型，目前取值Linux， Windows，Other

	OsType ImageInfoOsType `json:"__os_type"`
	// 操作系统具体版本

	OsVersion *string `json:"__os_version,omitempty"`
	// 镜像平台分类

	Platform *ImageInfoPlatform `json:"__platform,omitempty"`
	// 市场镜像的产品ID

	Productcode *string `json:"__productcode,omitempty"`
	// 镜像来源表示该镜像支持密集存储。如果镜像支持密集存储性能，则值为true，否则无需增加该属性。

	SupportDiskintensive *string `json:"__support_diskintensive,omitempty"`
	// 表示该镜像支持高计算性能。如果镜像支持高计算性能，则值为true，否则无需增加该属性。

	SupportHighperformance *string `json:"__support_highperformance,omitempty"`
	// 如果镜像支持KVM，取值为true，否则无需增加该属性。

	SupportKvm *string `json:"__support_kvm,omitempty"`
	// 表示该镜像是支持KVM虚拟化平台下的GPU类型，如果不支持KVM虚拟机下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。

	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty"`
	// 如果镜像支持KVM虚拟化下Infiniband网卡类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”属性不共存。

	SupportKvmInfiniband *string `json:"__support_kvm_infiniband,omitempty"`
	// 表示该镜像支持超大内存。如果镜像支持超大内存，取值为true，否则无需增加该属性。

	SupportLargememory *string `json:"__support_largememory,omitempty"`
	// 如果镜像支持XEN，取值为true，否则无需增加该属性。

	SupportXen *string `json:"__support_xen,omitempty"`
	// 表示该镜像是支持XEN虚拟化平台下的GPU优化类型，如果不支持XEN虚拟化下GPU类型，无需添加该属性            。该属性与“__support_xen”和“__support_kvm”属性不共存。

	SupportXenGpuType *string `json:"__support_xen_gpu_type,omitempty"`
	// 如果镜像支持XEN虚拟化下HANA类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”             和“__support_kvm”属性不共存。

	SupportXenHana *string `json:"__support_xen_hana,omitempty"`
	// 表示当前镜像是否支持发布为市场镜像,true表示支持,false 表示不支持

	SystemSupportMarket *bool `json:"__system_support_market,omitempty"`
	// 目前暂时不使用

	Checksum *string `json:"checksum,omitempty"`
	// 容器类型

	ContainerFormat string `json:"container_format"`
	// 创建时间。格式为UTC时间

	CreatedAt string `json:"created_at"`
	// 镜像的格式，目前支持vhd，zvhd、raw，qcow2。默认值是vhd

	DiskFormat *string `json:"disk_format,omitempty"`
	// 表示当前镜像所属的企业项目。取值为0或无该值，表示属于default企业项目，取值为UUID，表示属于该UUID对应的企业项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 镜像文件下载和上传链接

	File *string `json:"file,omitempty"`
	// 镜像ID

	Id string `json:"id"`
	// 镜像运行需要的最小磁盘容量，单位为GB。取值为40～1024GB。

	MinDisk int32 `json:"min_disk"`
	// 镜像运行需要的最小内存，单位为MB。参数取值依据弹性云服务器的规格限制，默认设置为0

	MinRam int32 `json:"min_ram"`
	// 镜像名称。 名称的首尾字母不能为空格。 名称的长度至为1～128位。 名称包含以下4种字符： 大写字母 小写字母 数字 特殊字符包含-、.、_、空格和中文。

	Name string `json:"name"`
	// 镜像属于哪个租户

	Owner string `json:"owner"`
	// 是否是受保护的，受保护的镜像不允许删除。取值为true或false

	Protected bool `json:"protected"`
	// 镜像视图

	Schema *string `json:"schema,omitempty"`
	// 镜像链接信息

	Self string `json:"self"`
	// 目前暂时不使用

	Size *int32 `json:"size,omitempty"`
	// 镜像状态。取值如下：queued：表示镜像元数据已经创建成功，等待 上传镜像文件。saving：表示镜像 正在上传文件到后端存储。deleted：表示镜像已经删除。killed：表示镜像上传错误。active：表示镜像可以正常使用

	Status ImageInfoStatus `json:"status"`
	// 镜像标签列表

	Tags []string `json:"tags"`
	// 更新时间。格式为UTC时间

	UpdatedAt string `json:"updated_at"`
	// 镜像使用环境类型：FusionCompute，Ironic，DataImage。如果弹性云服务器镜像，则取值为FusionCompute，如果是数据卷镜像则取Dat            aImage，如果是裸金属服务器镜像，则取值是Ironic

	VirtualEnvType ImageInfoVirtualEnvType `json:"virtual_env_type"`
	// 目前暂时不使用

	VirtualSize *int32 `json:"virtual_size,omitempty"`
	// 是否被其他租户可见，取值为private或public

	Visibility ImageInfoVisibility `json:"visibility"`
	// 表示当前镜像支持CloudInit密码/密钥注入方式，建议设置为\"true\"或者\"false\"。 如果取值为\"true\"，表示该镜像不支持CloudInit注入密码/密钥，其他取值时表示支持CloudInit注入密钥/密码。

	SupportFcInject *ImageInfoSupportFcInject `json:"__support_fc_inject,omitempty"`
	// 云服务器的启动方式。目前支持： bios：表示bios引导启动。 uefi：表示uefi引导启动。

	HwFirmwareType *ImageInfoHwFirmwareType `json:"hw_firmware_type,omitempty"`
	// 是否是ARM架构类型的镜像，取值为“true”或者“false”。

	SupportArm *ImageInfoSupportArm `json:"__support_arm,omitempty"`
	// 镜像支持的最大内存，单位为MB。取值可以参考云服务器规格限制，一般不设置。

	MaxRam *string `json:"max_ram,omitempty"`
	// 加密镜像所使用的密钥ID。

	SystemCmkid *string `json:"__system__cmkid,omitempty"`
	// 镜像附加属性。该属性采用JSON格式来标识镜像支持的高级特性清单。

	OsFeatureList *string `json:"__os_feature_list,omitempty"`
	// 收费镜像标识。

	AccountCode *string `json:"__account_code,omitempty"`
	// 镜像是否支持网卡多队列。取值为“true”或者“false”。

	HwVifMultiqueueEnabled *string `json:"hw_vif_multiqueue_enabled,omitempty"`
	// 表示当前市场镜像是否下架。true：已下架 false：未下架

	IsOffshelved *string `json:"__is_offshelved,omitempty"`
	// 镜像是否支持延迟加载。取值为“True”或“False”。

	Lazyloading *string `json:"__lazyloading,omitempty"`
	// 表示当前镜像来源是从外部导入。取值：file。

	RootOrigin *string `json:"__root_origin,omitempty"`
	// 表示当前镜像对应云服务器的系统盘插槽位置。目前暂时不用

	SequenceNum *string `json:"__sequence_num,omitempty"`
	// 镜像状态变为正常的时间。

	ActiveAt string `json:"active_at"`
	// 镜像是否支持企业主机安全或主机监控。 hss：企业主机安全 ces：主机监控

	SupportAgentList *string `json:"__support_agent_list,omitempty"`
	// 是否是AMD架构类型的镜像。取值为“true”或者“false”。

	SupportAmd *string `json:"__support_amd,omitempty"`
}

func (o ImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageInfo struct{}"
	}

	return strings.Join([]string{"ImageInfo", string(data)}, " ")
}

type ImageInfoImageSourceType struct {
	value string
}

type ImageInfoImageSourceTypeEnum struct {
	UDS   ImageInfoImageSourceType
	SWIFT ImageInfoImageSourceType
}

func GetImageInfoImageSourceTypeEnum() ImageInfoImageSourceTypeEnum {
	return ImageInfoImageSourceTypeEnum{
		UDS: ImageInfoImageSourceType{
			value: "uds",
		},
		SWIFT: ImageInfoImageSourceType{
			value: "swift",
		},
	}
}

func (c ImageInfoImageSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoImageSourceType) UnmarshalJSON(b []byte) error {
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

type ImageInfoImagetype struct {
	value string
}

type ImageInfoImagetypeEnum struct {
	GOLD    ImageInfoImagetype
	PRIVATE ImageInfoImagetype
	SHARED  ImageInfoImagetype
}

func GetImageInfoImagetypeEnum() ImageInfoImagetypeEnum {
	return ImageInfoImagetypeEnum{
		GOLD: ImageInfoImagetype{
			value: "gold",
		},
		PRIVATE: ImageInfoImagetype{
			value: "private",
		},
		SHARED: ImageInfoImagetype{
			value: "shared",
		},
	}
}

func (c ImageInfoImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoImagetype) UnmarshalJSON(b []byte) error {
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

type ImageInfoIsregistered struct {
	value string
}

type ImageInfoIsregisteredEnum struct {
	TRUE  ImageInfoIsregistered
	FALSE ImageInfoIsregistered
}

func GetImageInfoIsregisteredEnum() ImageInfoIsregisteredEnum {
	return ImageInfoIsregisteredEnum{
		TRUE: ImageInfoIsregistered{
			value: "true",
		},
		FALSE: ImageInfoIsregistered{
			value: "false",
		},
	}
}

func (c ImageInfoIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoIsregistered) UnmarshalJSON(b []byte) error {
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

type ImageInfoOsBit struct {
	value string
}

type ImageInfoOsBitEnum struct {
	E_32 ImageInfoOsBit
	E_64 ImageInfoOsBit
}

func GetImageInfoOsBitEnum() ImageInfoOsBitEnum {
	return ImageInfoOsBitEnum{
		E_32: ImageInfoOsBit{
			value: "32",
		},
		E_64: ImageInfoOsBit{
			value: "64",
		},
	}
}

func (c ImageInfoOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoOsBit) UnmarshalJSON(b []byte) error {
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

type ImageInfoOsType struct {
	value string
}

type ImageInfoOsTypeEnum struct {
	LINUX   ImageInfoOsType
	WINDOWS ImageInfoOsType
	OTHER   ImageInfoOsType
}

func GetImageInfoOsTypeEnum() ImageInfoOsTypeEnum {
	return ImageInfoOsTypeEnum{
		LINUX: ImageInfoOsType{
			value: "Linux",
		},
		WINDOWS: ImageInfoOsType{
			value: "Windows",
		},
		OTHER: ImageInfoOsType{
			value: "Other",
		},
	}
}

func (c ImageInfoOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoOsType) UnmarshalJSON(b []byte) error {
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

type ImageInfoPlatform struct {
	value string
}

type ImageInfoPlatformEnum struct {
	WINDOWS      ImageInfoPlatform
	UBUNTU       ImageInfoPlatform
	RED_HAT      ImageInfoPlatform
	SUSE         ImageInfoPlatform
	CENT_OS      ImageInfoPlatform
	DEBIAN       ImageInfoPlatform
	OPEN_SUSE    ImageInfoPlatform
	ORACLE_LINUX ImageInfoPlatform
	FEDORA       ImageInfoPlatform
	OTHER        ImageInfoPlatform
	CORE_OS      ImageInfoPlatform
	EULER_OS     ImageInfoPlatform
}

func GetImageInfoPlatformEnum() ImageInfoPlatformEnum {
	return ImageInfoPlatformEnum{
		WINDOWS: ImageInfoPlatform{
			value: "Windows",
		},
		UBUNTU: ImageInfoPlatform{
			value: "Ubuntu",
		},
		RED_HAT: ImageInfoPlatform{
			value: "RedHat",
		},
		SUSE: ImageInfoPlatform{
			value: "SUSE",
		},
		CENT_OS: ImageInfoPlatform{
			value: "CentOS",
		},
		DEBIAN: ImageInfoPlatform{
			value: "Debian",
		},
		OPEN_SUSE: ImageInfoPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: ImageInfoPlatform{
			value: "Oracle Linux",
		},
		FEDORA: ImageInfoPlatform{
			value: "Fedora",
		},
		OTHER: ImageInfoPlatform{
			value: "Other",
		},
		CORE_OS: ImageInfoPlatform{
			value: "CoreOS",
		},
		EULER_OS: ImageInfoPlatform{
			value: "EulerOS",
		},
	}
}

func (c ImageInfoPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoPlatform) UnmarshalJSON(b []byte) error {
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

type ImageInfoStatus struct {
	value string
}

type ImageInfoStatusEnum struct {
	QUEUED  ImageInfoStatus
	SAVING  ImageInfoStatus
	DELETED ImageInfoStatus
	KILLED  ImageInfoStatus
	ACTIVE  ImageInfoStatus
}

func GetImageInfoStatusEnum() ImageInfoStatusEnum {
	return ImageInfoStatusEnum{
		QUEUED: ImageInfoStatus{
			value: "queued",
		},
		SAVING: ImageInfoStatus{
			value: "saving",
		},
		DELETED: ImageInfoStatus{
			value: "deleted",
		},
		KILLED: ImageInfoStatus{
			value: "killed",
		},
		ACTIVE: ImageInfoStatus{
			value: "active",
		},
	}
}

func (c ImageInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoStatus) UnmarshalJSON(b []byte) error {
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

type ImageInfoVirtualEnvType struct {
	value string
}

type ImageInfoVirtualEnvTypeEnum struct {
	FUSION_COMPUTE ImageInfoVirtualEnvType
	IRONIC         ImageInfoVirtualEnvType
	DATA_IMAGE     ImageInfoVirtualEnvType
}

func GetImageInfoVirtualEnvTypeEnum() ImageInfoVirtualEnvTypeEnum {
	return ImageInfoVirtualEnvTypeEnum{
		FUSION_COMPUTE: ImageInfoVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: ImageInfoVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: ImageInfoVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c ImageInfoVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type ImageInfoVisibility struct {
	value string
}

type ImageInfoVisibilityEnum struct {
	PRIVATE ImageInfoVisibility
	PUBLIC  ImageInfoVisibility
}

func GetImageInfoVisibilityEnum() ImageInfoVisibilityEnum {
	return ImageInfoVisibilityEnum{
		PRIVATE: ImageInfoVisibility{
			value: "private",
		},
		PUBLIC: ImageInfoVisibility{
			value: "public",
		},
	}
}

func (c ImageInfoVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoVisibility) UnmarshalJSON(b []byte) error {
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

type ImageInfoSupportFcInject struct {
	value string
}

type ImageInfoSupportFcInjectEnum struct {
	TRUE  ImageInfoSupportFcInject
	FALSE ImageInfoSupportFcInject
}

func GetImageInfoSupportFcInjectEnum() ImageInfoSupportFcInjectEnum {
	return ImageInfoSupportFcInjectEnum{
		TRUE: ImageInfoSupportFcInject{
			value: "true",
		},
		FALSE: ImageInfoSupportFcInject{
			value: "false",
		},
	}
}

func (c ImageInfoSupportFcInject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoSupportFcInject) UnmarshalJSON(b []byte) error {
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

type ImageInfoHwFirmwareType struct {
	value string
}

type ImageInfoHwFirmwareTypeEnum struct {
	BIOS ImageInfoHwFirmwareType
	UEFI ImageInfoHwFirmwareType
}

func GetImageInfoHwFirmwareTypeEnum() ImageInfoHwFirmwareTypeEnum {
	return ImageInfoHwFirmwareTypeEnum{
		BIOS: ImageInfoHwFirmwareType{
			value: "bios",
		},
		UEFI: ImageInfoHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c ImageInfoHwFirmwareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoHwFirmwareType) UnmarshalJSON(b []byte) error {
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

type ImageInfoSupportArm struct {
	value string
}

type ImageInfoSupportArmEnum struct {
	TRUE  ImageInfoSupportArm
	FALSE ImageInfoSupportArm
}

func GetImageInfoSupportArmEnum() ImageInfoSupportArmEnum {
	return ImageInfoSupportArmEnum{
		TRUE: ImageInfoSupportArm{
			value: "true",
		},
		FALSE: ImageInfoSupportArm{
			value: "false",
		},
	}
}

func (c ImageInfoSupportArm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoSupportArm) UnmarshalJSON(b []byte) error {
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
