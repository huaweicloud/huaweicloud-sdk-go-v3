package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type GlanceShowImageResponse struct {

	// 备份ID。如果是备份创建的镜像，则填写为备份的ID，否则为空。
	BackupId *string `json:"__backup_id,omitempty" xml:"__backup_id"`

	// 镜像来源。公共镜像为空。
	DataOrigin *string `json:"__data_origin,omitempty" xml:"__data_origin"`

	// 镜像描述信息。
	Description *string `json:"__description,omitempty" xml:"__description"`

	// 镜像文件的大小，单位为字节。目前取值为大于0的字符串。
	ImageSize *string `json:"__image_size,omitempty" xml:"__image_size"`

	// 镜像后端存储类型，目前只支持uds
	ImageSourceType *GlanceShowImageResponseImageSourceType `json:"__image_source_type,omitempty" xml:"__image_source_type"`

	// 镜像类型，目前支持以下类型：公共镜像：gold私有镜像：private共享镜像：shared
	Imagetype *GlanceShowImageResponseImagetype `json:"__imagetype,omitempty" xml:"__imagetype"`

	// 是否是注册过的镜像，取值为“true”或者“false”。
	Isregistered *GlanceShowImageResponseIsregistered `json:"__isregistered,omitempty" xml:"__isregistered"`

	// 父镜像ID。公共镜像或通过文件创建的私有镜像，取值为空。
	Originalimagename *string `json:"__originalimagename,omitempty" xml:"__originalimagename"`

	// 操作系统位数，一般取值为“32”或者“64”。
	OsBit *GlanceShowImageResponseOsBit `json:"__os_bit,omitempty" xml:"__os_bit"`

	// 操作系统类型，目前取值Linux， Windows，Other。
	OsType *GlanceShowImageResponseOsType `json:"__os_type,omitempty" xml:"__os_type"`

	// 操作系统具体版本。
	OsVersion *string `json:"__os_version,omitempty" xml:"__os_version"`

	// 镜像平台分类，取值为Windows，Ubuntu，RedHat，SUSE，CentOS，Debian，OpenSUSE, Oracle Linux，Fedora，Other，CoreOS和EulerOS。
	Platform *GlanceShowImageResponsePlatform `json:"__platform,omitempty" xml:"__platform"`

	// 市场镜像的产品ID。
	Productcode *string `json:"__productcode,omitempty" xml:"__productcode"`

	// 表示该镜像支持密集存储。如果镜像支持密集存储性能，则值为true，否则无需增加该属性。
	SupportDiskintensive *string `json:"__support_diskintensive,omitempty" xml:"__support_diskintensive"`

	// 表示该镜像支持高计算性能。如果镜像支持高计算性能，则值为true，否则无需增加该属性。
	SupportHighperformance *string `json:"__support_highperformance,omitempty" xml:"__support_highperformance"`

	// 如果镜像支持KVM，取值为true，否则无需增加该属性。
	SupportKvm *string `json:"__support_kvm,omitempty" xml:"__support_kvm"`

	// 表示该镜像是支持KVM虚拟化平台下的GPU类型，如果不支持KVM虚拟机下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty" xml:"__support_kvm_gpu_type"`

	// 如果镜像支持KVM虚拟化下Infiniband网卡类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”属性不共存。
	SupportKvmInfiniband *string `json:"__support_kvm_infiniband,omitempty" xml:"__support_kvm_infiniband"`

	// 表示该镜像支持超大内存。如果镜像支持超大内存，取值为true，否则无需增加该属性
	SupportLargememory *string `json:"__support_largememory,omitempty" xml:"__support_largememory"`

	// 如果镜像支持XEN，取值为true，否则无需增加该属性。
	SupportXen *string `json:"__support_xen,omitempty" xml:"__support_xen"`

	// 表示该镜像是支持XEN虚拟化平台下的GPU优化类型，取值参考8.10-表 镜像支持的GPU类型说明。镜像操作系统类型请参考8.10-表 镜像类型。如果不支持XEN虚拟化下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenGpuType *string `json:"__support_xen_gpu_type,omitempty" xml:"__support_xen_gpu_type"`

	// 如果镜像支持XEN虚拟化下HANA类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenHana *string `json:"__support_xen_hana,omitempty" xml:"__support_xen_hana"`

	// 目前暂时不使用。
	Checksum *string `json:"checksum,omitempty" xml:"checksum"`

	// 容器类型。
	ContainerFormat *string `json:"container_format,omitempty" xml:"container_format"`

	// 创建时间。格式为UTC时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 镜像的格式，目前支持vhd，zvhd、raw，qcow2,zvhd2。默认值是vhd。
	DiskFormat *GlanceShowImageResponseDiskFormat `json:"disk_format,omitempty" xml:"disk_format"`

	// 镜像文件下载和上传链接。
	File *string `json:"file,omitempty" xml:"file"`

	// 镜像ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 镜像运行需要的最小磁盘容量，单位为GB
	MinDisk *int32 `json:"min_disk,omitempty" xml:"min_disk"`

	// 镜像运行最小内存，单位为MB。
	MinRam *int32 `json:"min_ram,omitempty" xml:"min_ram"`

	// 镜像名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 镜像属于哪个租户。
	Owner *string `json:"owner,omitempty" xml:"owner"`

	// 是否是受保护的，受保护的镜像不允许删除。取值为true或false。
	Protected *bool `json:"protected,omitempty" xml:"protected"`

	// 镜像视图。
	Schema *string `json:"schema,omitempty" xml:"schema"`

	// 镜像链接信息。
	Self *string `json:"self,omitempty" xml:"self"`

	// 目前暂时不使用。
	Size *int64 `json:"size,omitempty" xml:"size"`

	// 镜像状态。取值如下：queued：表示镜像元数据已经创建成功，等待上传镜像文件。saving：表示镜像正在上传文件到后端存储。deleted：表示镜像已经删除。killed：表示镜像上传错误。active：表示镜像可以正常使用。
	Status *GlanceShowImageResponseStatus `json:"status,omitempty" xml:"status"`

	// 镜像标签列表，提供用户可以自定义管理私有镜像的能力。用户可以通过镜像标签接口为每个镜像增加不同的标签，在查询接口中可以根据标签进行过滤。
	Tags *[]string `json:"tags,omitempty" xml:"tags"`

	// 更新时间。格式为UTC时间。
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 镜像使用环境类型：FusionCompute，Ironic，DataImage。
	VirtualEnvType *GlanceShowImageResponseVirtualEnvType `json:"virtual_env_type,omitempty" xml:"virtual_env_type"`

	// 目前暂时不使用。
	VirtualSize *int32 `json:"virtual_size,omitempty" xml:"virtual_size"`

	// 是否被其他租户可见，取值如下：private：私有镜像public：公共镜像shared：共享镜像
	Visibility *GlanceShowImageResponseVisibility `json:"visibility,omitempty" xml:"visibility"`

	// 表示当前镜像支持CloudInit密码/密钥注入方式，建议设置为\"true\"或者\"false\"。如果取值为\"true\"，表示该镜像不支持CloudInit注入密码/密钥，其他取值时表示支持CloudInit注入密钥/密码。
	SupportFcInject *GlanceShowImageResponseSupportFcInject `json:"__support_fc_inject,omitempty" xml:"__support_fc_inject"`

	// 表示当前镜像所属的企业项目。 取值为0或无该值，表示属于default企业项目。 取值为UUID，表示属于该UUID对应的企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 云主机云服务器的启动方式。目前支持： bios：表示bios引导启动。 uefi：表示uefi引导启动。
	HwFirmwareType *GlanceShowImageResponseHwFirmwareType `json:"hw_firmware_type,omitempty" xml:"hw_firmware_type"`

	// 是否为ARM架构类型的镜像，取值为“true”或者“false”。
	SupportArm *GlanceShowImageResponseSupportArm `json:"__support_arm,omitempty" xml:"__support_arm"`

	// 表示当前市场镜像是否下架。 true：已下架 false：未下架
	IsOffshelved *GlanceShowImageResponseIsOffshelved `json:"__is_offshelved,omitempty" xml:"__is_offshelved"`

	// 镜像是否支持延迟加载。取值为True或False。
	Lazyloading *string `json:"__lazyloading,omitempty" xml:"__lazyloading"`

	// 镜像附加属性。该属性采用JSON格式来标识镜像支持的高级特性清单。
	OsFeatureList *string `json:"__os_feature_list,omitempty" xml:"__os_feature_list"`

	// 表示当前镜像来源是从外部导入。取值：file。
	RootOrigin *string `json:"__root_origin,omitempty" xml:"__root_origin"`

	// 目前暂时不用
	SequenceNum *string `json:"__sequence_num,omitempty" xml:"__sequence_num"`

	// 镜像是否支持企业主机安全或主机监控。 hss：企业主机安全 ces：主机监控
	SupportAgentList *string `json:"__support_agent_list,omitempty" xml:"__support_agent_list"`

	// 加密镜像所使用的密钥ID。
	SystemCmkid *string `json:"__system__cmkid,omitempty" xml:"__system__cmkid"`

	// 镜像状态变为正常的时间。
	ActiveAt *string `json:"active_at,omitempty" xml:"active_at"`

	// 镜像是否支持网卡多队列。取值为true或者false。
	HwVifMultiqueueEnabled *string `json:"hw_vif_multiqueue_enabled,omitempty" xml:"hw_vif_multiqueue_enabled"`

	// 镜像支持的最大内存，单位为MB。取值可以参考云服务器规格限制，一般不设置。
	MaxRam *string `json:"max_ram,omitempty" xml:"max_ram"`

	// 镜像的存储位置。
	ImageLocation *string `json:"__image_location,omitempty" xml:"__image_location"`

	// 是否完成了初始化配置。取值为true或false
	IsConfigInit *string `json:"__is_config_init,omitempty" xml:"__is_config_init"`

	// 收费镜像标识。
	AccountCode *string `json:"__account_code,omitempty" xml:"__account_code"`

	// 是否是AMD架构类型的镜像。取值为“true”或者“false”。
	SupportAmd     *string `json:"__support_amd,omitempty" xml:"__support_amd"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceShowImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageResponse struct{}"
	}

	return strings.Join([]string{"GlanceShowImageResponse", string(data)}, " ")
}

type GlanceShowImageResponseImageSourceType struct {
	value string
}

type GlanceShowImageResponseImageSourceTypeEnum struct {
	UDS   GlanceShowImageResponseImageSourceType
	SWIFT GlanceShowImageResponseImageSourceType
}

func GetGlanceShowImageResponseImageSourceTypeEnum() GlanceShowImageResponseImageSourceTypeEnum {
	return GlanceShowImageResponseImageSourceTypeEnum{
		UDS: GlanceShowImageResponseImageSourceType{
			value: "uds",
		},
		SWIFT: GlanceShowImageResponseImageSourceType{
			value: "swift",
		},
	}
}

func (c GlanceShowImageResponseImageSourceType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseImageSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseImageSourceType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseImagetype struct {
	value string
}

type GlanceShowImageResponseImagetypeEnum struct {
	GOLD    GlanceShowImageResponseImagetype
	PRIVATE GlanceShowImageResponseImagetype
	SHARED  GlanceShowImageResponseImagetype
}

func GetGlanceShowImageResponseImagetypeEnum() GlanceShowImageResponseImagetypeEnum {
	return GlanceShowImageResponseImagetypeEnum{
		GOLD: GlanceShowImageResponseImagetype{
			value: "gold",
		},
		PRIVATE: GlanceShowImageResponseImagetype{
			value: "private",
		},
		SHARED: GlanceShowImageResponseImagetype{
			value: "shared",
		},
	}
}

func (c GlanceShowImageResponseImagetype) Value() string {
	return c.value
}

func (c GlanceShowImageResponseImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseImagetype) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseIsregistered struct {
	value string
}

type GlanceShowImageResponseIsregisteredEnum struct {
	TRUE  GlanceShowImageResponseIsregistered
	FALSE GlanceShowImageResponseIsregistered
}

func GetGlanceShowImageResponseIsregisteredEnum() GlanceShowImageResponseIsregisteredEnum {
	return GlanceShowImageResponseIsregisteredEnum{
		TRUE: GlanceShowImageResponseIsregistered{
			value: "true",
		},
		FALSE: GlanceShowImageResponseIsregistered{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseIsregistered) Value() string {
	return c.value
}

func (c GlanceShowImageResponseIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseIsregistered) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseOsBit struct {
	value string
}

type GlanceShowImageResponseOsBitEnum struct {
	E_32 GlanceShowImageResponseOsBit
	E_64 GlanceShowImageResponseOsBit
}

func GetGlanceShowImageResponseOsBitEnum() GlanceShowImageResponseOsBitEnum {
	return GlanceShowImageResponseOsBitEnum{
		E_32: GlanceShowImageResponseOsBit{
			value: "32",
		},
		E_64: GlanceShowImageResponseOsBit{
			value: "64",
		},
	}
}

func (c GlanceShowImageResponseOsBit) Value() string {
	return c.value
}

func (c GlanceShowImageResponseOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseOsType struct {
	value string
}

type GlanceShowImageResponseOsTypeEnum struct {
	LINUX   GlanceShowImageResponseOsType
	WINDOWS GlanceShowImageResponseOsType
	OTHER   GlanceShowImageResponseOsType
}

func GetGlanceShowImageResponseOsTypeEnum() GlanceShowImageResponseOsTypeEnum {
	return GlanceShowImageResponseOsTypeEnum{
		LINUX: GlanceShowImageResponseOsType{
			value: "Linux",
		},
		WINDOWS: GlanceShowImageResponseOsType{
			value: "Windows",
		},
		OTHER: GlanceShowImageResponseOsType{
			value: "Other",
		},
	}
}

func (c GlanceShowImageResponseOsType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseOsType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponsePlatform struct {
	value string
}

type GlanceShowImageResponsePlatformEnum struct {
	WINDOWS      GlanceShowImageResponsePlatform
	UBUNTU       GlanceShowImageResponsePlatform
	RED_HAT      GlanceShowImageResponsePlatform
	SUSE         GlanceShowImageResponsePlatform
	CENT_OS      GlanceShowImageResponsePlatform
	DEBIAN       GlanceShowImageResponsePlatform
	OPEN_SUSE    GlanceShowImageResponsePlatform
	ORACLE_LINUX GlanceShowImageResponsePlatform
	FEDORA       GlanceShowImageResponsePlatform
	OTHER        GlanceShowImageResponsePlatform
	CORE_OS      GlanceShowImageResponsePlatform
	EULER_OS     GlanceShowImageResponsePlatform
}

func GetGlanceShowImageResponsePlatformEnum() GlanceShowImageResponsePlatformEnum {
	return GlanceShowImageResponsePlatformEnum{
		WINDOWS: GlanceShowImageResponsePlatform{
			value: "Windows",
		},
		UBUNTU: GlanceShowImageResponsePlatform{
			value: "Ubuntu",
		},
		RED_HAT: GlanceShowImageResponsePlatform{
			value: "RedHat",
		},
		SUSE: GlanceShowImageResponsePlatform{
			value: "SUSE",
		},
		CENT_OS: GlanceShowImageResponsePlatform{
			value: "CentOS",
		},
		DEBIAN: GlanceShowImageResponsePlatform{
			value: "Debian",
		},
		OPEN_SUSE: GlanceShowImageResponsePlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: GlanceShowImageResponsePlatform{
			value: "OracleLinux",
		},
		FEDORA: GlanceShowImageResponsePlatform{
			value: "Fedora",
		},
		OTHER: GlanceShowImageResponsePlatform{
			value: "Other",
		},
		CORE_OS: GlanceShowImageResponsePlatform{
			value: "CoreOS",
		},
		EULER_OS: GlanceShowImageResponsePlatform{
			value: "EulerOS",
		},
	}
}

func (c GlanceShowImageResponsePlatform) Value() string {
	return c.value
}

func (c GlanceShowImageResponsePlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponsePlatform) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseDiskFormat struct {
	value string
}

type GlanceShowImageResponseDiskFormatEnum struct {
	VHD   GlanceShowImageResponseDiskFormat
	ZVHD  GlanceShowImageResponseDiskFormat
	RAW   GlanceShowImageResponseDiskFormat
	QCOW2 GlanceShowImageResponseDiskFormat
	ZVHD2 GlanceShowImageResponseDiskFormat
}

func GetGlanceShowImageResponseDiskFormatEnum() GlanceShowImageResponseDiskFormatEnum {
	return GlanceShowImageResponseDiskFormatEnum{
		VHD: GlanceShowImageResponseDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceShowImageResponseDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceShowImageResponseDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceShowImageResponseDiskFormat{
			value: "qcow2",
		},
		ZVHD2: GlanceShowImageResponseDiskFormat{
			value: "zvhd2",
		},
	}
}

func (c GlanceShowImageResponseDiskFormat) Value() string {
	return c.value
}

func (c GlanceShowImageResponseDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseStatus struct {
	value string
}

type GlanceShowImageResponseStatusEnum struct {
	QUEUED  GlanceShowImageResponseStatus
	SAVING  GlanceShowImageResponseStatus
	DELETED GlanceShowImageResponseStatus
	KILLED  GlanceShowImageResponseStatus
	ACTIVE  GlanceShowImageResponseStatus
}

func GetGlanceShowImageResponseStatusEnum() GlanceShowImageResponseStatusEnum {
	return GlanceShowImageResponseStatusEnum{
		QUEUED: GlanceShowImageResponseStatus{
			value: "queued",
		},
		SAVING: GlanceShowImageResponseStatus{
			value: "saving",
		},
		DELETED: GlanceShowImageResponseStatus{
			value: "deleted",
		},
		KILLED: GlanceShowImageResponseStatus{
			value: "killed",
		},
		ACTIVE: GlanceShowImageResponseStatus{
			value: "active",
		},
	}
}

func (c GlanceShowImageResponseStatus) Value() string {
	return c.value
}

func (c GlanceShowImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseVirtualEnvType struct {
	value string
}

type GlanceShowImageResponseVirtualEnvTypeEnum struct {
	FUSION_COMPUTE GlanceShowImageResponseVirtualEnvType
	IRONIC         GlanceShowImageResponseVirtualEnvType
	DATA_IMAGE     GlanceShowImageResponseVirtualEnvType
}

func GetGlanceShowImageResponseVirtualEnvTypeEnum() GlanceShowImageResponseVirtualEnvTypeEnum {
	return GlanceShowImageResponseVirtualEnvTypeEnum{
		FUSION_COMPUTE: GlanceShowImageResponseVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: GlanceShowImageResponseVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: GlanceShowImageResponseVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c GlanceShowImageResponseVirtualEnvType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseVisibility struct {
	value string
}

type GlanceShowImageResponseVisibilityEnum struct {
	PRIVATE GlanceShowImageResponseVisibility
	PUBLIC  GlanceShowImageResponseVisibility
	SHARED  GlanceShowImageResponseVisibility
}

func GetGlanceShowImageResponseVisibilityEnum() GlanceShowImageResponseVisibilityEnum {
	return GlanceShowImageResponseVisibilityEnum{
		PRIVATE: GlanceShowImageResponseVisibility{
			value: "private",
		},
		PUBLIC: GlanceShowImageResponseVisibility{
			value: "public",
		},
		SHARED: GlanceShowImageResponseVisibility{
			value: "shared",
		},
	}
}

func (c GlanceShowImageResponseVisibility) Value() string {
	return c.value
}

func (c GlanceShowImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseVisibility) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseSupportFcInject struct {
	value string
}

type GlanceShowImageResponseSupportFcInjectEnum struct {
	TRUE  GlanceShowImageResponseSupportFcInject
	FALSE GlanceShowImageResponseSupportFcInject
}

func GetGlanceShowImageResponseSupportFcInjectEnum() GlanceShowImageResponseSupportFcInjectEnum {
	return GlanceShowImageResponseSupportFcInjectEnum{
		TRUE: GlanceShowImageResponseSupportFcInject{
			value: "true",
		},
		FALSE: GlanceShowImageResponseSupportFcInject{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseSupportFcInject) Value() string {
	return c.value
}

func (c GlanceShowImageResponseSupportFcInject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseSupportFcInject) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseHwFirmwareType struct {
	value string
}

type GlanceShowImageResponseHwFirmwareTypeEnum struct {
	BIOS GlanceShowImageResponseHwFirmwareType
	UEFI GlanceShowImageResponseHwFirmwareType
}

func GetGlanceShowImageResponseHwFirmwareTypeEnum() GlanceShowImageResponseHwFirmwareTypeEnum {
	return GlanceShowImageResponseHwFirmwareTypeEnum{
		BIOS: GlanceShowImageResponseHwFirmwareType{
			value: "bios",
		},
		UEFI: GlanceShowImageResponseHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c GlanceShowImageResponseHwFirmwareType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseHwFirmwareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseHwFirmwareType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseSupportArm struct {
	value string
}

type GlanceShowImageResponseSupportArmEnum struct {
	TRUE  GlanceShowImageResponseSupportArm
	FALSE GlanceShowImageResponseSupportArm
}

func GetGlanceShowImageResponseSupportArmEnum() GlanceShowImageResponseSupportArmEnum {
	return GlanceShowImageResponseSupportArmEnum{
		TRUE: GlanceShowImageResponseSupportArm{
			value: "true",
		},
		FALSE: GlanceShowImageResponseSupportArm{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseSupportArm) Value() string {
	return c.value
}

func (c GlanceShowImageResponseSupportArm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseSupportArm) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseIsOffshelved struct {
	value string
}

type GlanceShowImageResponseIsOffshelvedEnum struct {
	TRUE  GlanceShowImageResponseIsOffshelved
	FALSE GlanceShowImageResponseIsOffshelved
}

func GetGlanceShowImageResponseIsOffshelvedEnum() GlanceShowImageResponseIsOffshelvedEnum {
	return GlanceShowImageResponseIsOffshelvedEnum{
		TRUE: GlanceShowImageResponseIsOffshelved{
			value: "true",
		},
		FALSE: GlanceShowImageResponseIsOffshelved{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseIsOffshelved) Value() string {
	return c.value
}

func (c GlanceShowImageResponseIsOffshelved) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseIsOffshelved) UnmarshalJSON(b []byte) error {
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
