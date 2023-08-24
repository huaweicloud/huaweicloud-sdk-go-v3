package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GlanceShowImageListResponseBody
type GlanceShowImageListResponseBody struct {

	// 备份ID。如果是备份创建的镜像，则填写为备份的ID，否则为空。
	BackupId *string `json:"__backup_id,omitempty"`

	// 镜像来源。公共镜像为空。
	DataOrigin *string `json:"__data_origin,omitempty"`

	// 镜像描述信息。
	Description *string `json:"__description,omitempty"`

	// 镜像文件的大小，单位为字节。目前取值为大于0的字符串。
	ImageSize string `json:"__image_size"`

	// 镜像后端存储类型，目前只支持uds
	ImageSourceType GlanceShowImageListResponseBodyImageSourceType `json:"__image_source_type"`

	// 镜像类型，目前支持以下类型：公共镜像：gold私有镜像：private共享镜像：shared 市场镜像：market
	Imagetype GlanceShowImageListResponseBodyImagetype `json:"__imagetype"`

	// 是否是注册过的镜像，取值为“true”或者“false”。
	Isregistered GlanceShowImageListResponseBodyIsregistered `json:"__isregistered"`

	// 父镜像ID。公共镜像或通过文件创建的私有镜像，取值为空。
	Originalimagename *string `json:"__originalimagename,omitempty"`

	// 操作系统位数，一般取值为“32”或者“64”。
	OsBit *GlanceShowImageListResponseBodyOsBit `json:"__os_bit,omitempty"`

	// 操作系统类型，目前取值Linux， Windows，Other。
	OsType GlanceShowImageListResponseBodyOsType `json:"__os_type"`

	// 操作系统具体版本。
	OsVersion *string `json:"__os_version,omitempty"`

	// 镜像平台分类，取值为Windows，Ubuntu，RedHat，SUSE，CentOS，Debian，OpenSUSE, Oracle Linux，Fedora，Other，CoreOS和EulerOS。
	Platform *GlanceShowImageListResponseBodyPlatform `json:"__platform,omitempty"`

	// 市场镜像的产品ID。
	Productcode *string `json:"__productcode,omitempty"`

	// 表示该镜像支持密集存储。如果镜像支持密集存储性能，则值为true，否则无需增加该属性。
	SupportDiskintensive *string `json:"__support_diskintensive,omitempty"`

	// 表示该镜像支持高计算性能。如果镜像支持高计算性能，则值为true，否则无需增加该属性。
	SupportHighperformance *string `json:"__support_highperformance,omitempty"`

	// 如果镜像支持KVM，取值为true，否则无需增加该属性。
	SupportKvm *string `json:"__support_kvm,omitempty"`

	// 表示该镜像是支持KVM虚拟化平台下的GPU类型，如果不支持KVM虚拟机下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty"`

	// 如果镜像支持KVM虚拟化下Infiniband网卡类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”属性不共存。
	SupportKvmInfiniband *string `json:"__support_kvm_infiniband,omitempty"`

	// 表示该镜像支持超大内存。如果镜像支持超大内存，取值为true，否则无需增加该属性
	SupportLargememory *string `json:"__support_largememory,omitempty"`

	// 如果镜像支持XEN，取值为true，否则无需增加该属性。
	SupportXen *string `json:"__support_xen,omitempty"`

	// 表示该镜像是支持XEN虚拟化平台下的GPU优化类型，取值参考8.10-表 镜像支持的GPU类型说明。镜像操作系统类型请参考8.10-表 镜像类型。如果不支持XEN虚拟化下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenGpuType *string `json:"__support_xen_gpu_type,omitempty"`

	// 如果镜像支持XEN虚拟化下HANA类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenHana *string `json:"__support_xen_hana,omitempty"`

	// 目前暂时不使用。
	Checksum *string `json:"checksum,omitempty"`

	// 容器类型。
	ContainerFormat string `json:"container_format"`

	// 创建时间。格式为UTC时间。
	CreatedAt string `json:"created_at"`

	// 镜像的格式，目前支持vhd，zvhd、raw，qcow2,zvhd2。默认值是vhd。
	DiskFormat GlanceShowImageListResponseBodyDiskFormat `json:"disk_format"`

	// 镜像文件下载和上传链接。
	File string `json:"file"`

	// 镜像ID。
	Id string `json:"id"`

	// 镜像运行需要的最小磁盘容量，单位为GB
	MinDisk int32 `json:"min_disk"`

	// 镜像运行最小内存，单位为MB。
	MinRam int32 `json:"min_ram"`

	// 镜像名称。
	Name string `json:"name"`

	// 镜像属于哪个租户。
	Owner string `json:"owner"`

	// 是否是受保护的，受保护的镜像不允许删除。取值为true或false。
	Protected bool `json:"protected"`

	// 镜像视图。
	Schema string `json:"schema"`

	// 镜像链接信息。
	Self string `json:"self"`

	// 目前暂时不使用。
	Size *int64 `json:"size,omitempty"`

	// 镜像状态。取值如下：queued：表示镜像元数据已经创建成功，等待上传镜像文件。saving：表示镜像正在上传文件到后端存储。deleted：表示镜像已经删除。killed：表示镜像上传错误。active：表示镜像可以正常使用。
	Status GlanceShowImageListResponseBodyStatus `json:"status"`

	// 镜像标签列表，提供用户可以自定义管理私有镜像的能力。用户可以通过镜像标签接口为每个镜像增加不同的标签，在查询接口中可以根据标签进行过滤。
	Tags []string `json:"tags"`

	// 更新时间。格式为UTC时间。
	UpdatedAt string `json:"updated_at"`

	// 镜像使用环境类型：FusionCompute，Ironic，DataImage。
	VirtualEnvType GlanceShowImageListResponseBodyVirtualEnvType `json:"virtual_env_type"`

	// 目前暂时不使用。
	VirtualSize *int32 `json:"virtual_size,omitempty"`

	// 是否被其他租户可见，取值如下：private：私有镜像public：公共镜像shared：共享镜像
	Visibility GlanceShowImageListResponseBodyVisibility `json:"visibility"`

	// 表示当前镜像支持CloudInit密码/密钥注入方式，建议设置为\"true\"或者\"false\"。如果取值为\"true\"，表示该镜像不支持CloudInit注入密码/密钥，其他取值时表示支持CloudInit注入密钥/密码。
	SupportFcInject *GlanceShowImageListResponseBodySupportFcInject `json:"__support_fc_inject,omitempty"`

	// 表示当前镜像所属的企业项目。 取值为0或无该值，表示属于default企业项目。 取值为UUID，表示属于该UUID对应的企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 表示当前镜像所属的企业项目。 取值为0或无该值，表示属于default企业项目。 取值为UUID，表示属于该UUID对应的企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。
	SysEnterpriseProjectId *string `json:"_sys_enterprise_project_id,omitempty"`

	// 云主机云服务器的启动方式。目前支持： bios：表示bios引导启动。 uefi：表示uefi引导启动。
	HwFirmwareType *GlanceShowImageListResponseBodyHwFirmwareType `json:"hw_firmware_type,omitempty"`

	// 是否为ARM架构类型的镜像，取值为“true”或者“false”。
	SupportArm *GlanceShowImageListResponseBodySupportArm `json:"__support_arm,omitempty"`

	// 表示当前市场镜像是否下架。 true：已下架 false：未下架
	IsOffshelved *GlanceShowImageListResponseBodyIsOffshelved `json:"__is_offshelved,omitempty"`

	// 镜像是否支持延迟加载。取值为True或False。
	Lazyloading *string `json:"__lazyloading,omitempty"`

	// 镜像附加属性。该属性采用JSON格式来标识镜像支持的高级特性清单。
	OsFeatureList *string `json:"__os_feature_list,omitempty"`

	// 表示当前镜像来源是从外部导入。取值：file。
	RootOrigin *string `json:"__root_origin,omitempty"`

	// 目前暂时不用
	SequenceNum *string `json:"__sequence_num,omitempty"`

	// 镜像是否支持企业主机安全或主机监控。 hss：企业主机安全 ces：主机监控
	SupportAgentList *string `json:"__support_agent_list,omitempty"`

	// 加密镜像所使用的密钥ID。
	SystemCmkid *string `json:"__system__cmkid,omitempty"`

	// 镜像状态变为正常的时间。
	ActiveAt *string `json:"active_at,omitempty"`

	// 镜像是否支持网卡多队列。取值为true或者false。
	HwVifMultiqueueEnabled *string `json:"hw_vif_multiqueue_enabled,omitempty"`

	// 镜像支持的最大内存，单位为MB。取值可以参考云服务器规格限制，一般不设置。
	MaxRam *string `json:"max_ram,omitempty"`

	// 镜像的存储位置。
	ImageLocation *string `json:"__image_location,omitempty"`

	// 是否完成了初始化配置。取值为true或false
	IsConfigInit *string `json:"__is_config_init,omitempty"`

	// 收费镜像标识。
	AccountCode *string `json:"__account_code,omitempty"`

	// 是否是AMD架构类型的镜像。取值为“true”或者“false”。
	SupportAmd *string `json:"__support_amd,omitempty"`
}

func (o GlanceShowImageListResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageListResponseBody struct{}"
	}

	return strings.Join([]string{"GlanceShowImageListResponseBody", string(data)}, " ")
}

type GlanceShowImageListResponseBodyImageSourceType struct {
	value string
}

type GlanceShowImageListResponseBodyImageSourceTypeEnum struct {
	UDS   GlanceShowImageListResponseBodyImageSourceType
	SWIFT GlanceShowImageListResponseBodyImageSourceType
}

func GetGlanceShowImageListResponseBodyImageSourceTypeEnum() GlanceShowImageListResponseBodyImageSourceTypeEnum {
	return GlanceShowImageListResponseBodyImageSourceTypeEnum{
		UDS: GlanceShowImageListResponseBodyImageSourceType{
			value: "uds",
		},
		SWIFT: GlanceShowImageListResponseBodyImageSourceType{
			value: "swift",
		},
	}
}

func (c GlanceShowImageListResponseBodyImageSourceType) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyImageSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyImageSourceType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyImagetype struct {
	value string
}

type GlanceShowImageListResponseBodyImagetypeEnum struct {
	GOLD    GlanceShowImageListResponseBodyImagetype
	PRIVATE GlanceShowImageListResponseBodyImagetype
	SHARED  GlanceShowImageListResponseBodyImagetype
	MARKET  GlanceShowImageListResponseBodyImagetype
}

func GetGlanceShowImageListResponseBodyImagetypeEnum() GlanceShowImageListResponseBodyImagetypeEnum {
	return GlanceShowImageListResponseBodyImagetypeEnum{
		GOLD: GlanceShowImageListResponseBodyImagetype{
			value: "gold",
		},
		PRIVATE: GlanceShowImageListResponseBodyImagetype{
			value: "private",
		},
		SHARED: GlanceShowImageListResponseBodyImagetype{
			value: "shared",
		},
		MARKET: GlanceShowImageListResponseBodyImagetype{
			value: "market",
		},
	}
}

func (c GlanceShowImageListResponseBodyImagetype) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyImagetype) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyIsregistered struct {
	value string
}

type GlanceShowImageListResponseBodyIsregisteredEnum struct {
	TRUE  GlanceShowImageListResponseBodyIsregistered
	FALSE GlanceShowImageListResponseBodyIsregistered
}

func GetGlanceShowImageListResponseBodyIsregisteredEnum() GlanceShowImageListResponseBodyIsregisteredEnum {
	return GlanceShowImageListResponseBodyIsregisteredEnum{
		TRUE: GlanceShowImageListResponseBodyIsregistered{
			value: "true",
		},
		FALSE: GlanceShowImageListResponseBodyIsregistered{
			value: "false",
		},
	}
}

func (c GlanceShowImageListResponseBodyIsregistered) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyIsregistered) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyOsBit struct {
	value string
}

type GlanceShowImageListResponseBodyOsBitEnum struct {
	E_32 GlanceShowImageListResponseBodyOsBit
	E_64 GlanceShowImageListResponseBodyOsBit
}

func GetGlanceShowImageListResponseBodyOsBitEnum() GlanceShowImageListResponseBodyOsBitEnum {
	return GlanceShowImageListResponseBodyOsBitEnum{
		E_32: GlanceShowImageListResponseBodyOsBit{
			value: "32",
		},
		E_64: GlanceShowImageListResponseBodyOsBit{
			value: "64",
		},
	}
}

func (c GlanceShowImageListResponseBodyOsBit) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyOsType struct {
	value string
}

type GlanceShowImageListResponseBodyOsTypeEnum struct {
	LINUX   GlanceShowImageListResponseBodyOsType
	WINDOWS GlanceShowImageListResponseBodyOsType
	OTHER   GlanceShowImageListResponseBodyOsType
}

func GetGlanceShowImageListResponseBodyOsTypeEnum() GlanceShowImageListResponseBodyOsTypeEnum {
	return GlanceShowImageListResponseBodyOsTypeEnum{
		LINUX: GlanceShowImageListResponseBodyOsType{
			value: "Linux",
		},
		WINDOWS: GlanceShowImageListResponseBodyOsType{
			value: "Windows",
		},
		OTHER: GlanceShowImageListResponseBodyOsType{
			value: "Other",
		},
	}
}

func (c GlanceShowImageListResponseBodyOsType) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyOsType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyPlatform struct {
	value string
}

type GlanceShowImageListResponseBodyPlatformEnum struct {
	WINDOWS      GlanceShowImageListResponseBodyPlatform
	UBUNTU       GlanceShowImageListResponseBodyPlatform
	RED_HAT      GlanceShowImageListResponseBodyPlatform
	SUSE         GlanceShowImageListResponseBodyPlatform
	CENT_OS      GlanceShowImageListResponseBodyPlatform
	DEBIAN       GlanceShowImageListResponseBodyPlatform
	OPEN_SUSE    GlanceShowImageListResponseBodyPlatform
	ORACLE_LINUX GlanceShowImageListResponseBodyPlatform
	FEDORA       GlanceShowImageListResponseBodyPlatform
	OTHER        GlanceShowImageListResponseBodyPlatform
	CORE_OS      GlanceShowImageListResponseBodyPlatform
	EULER_OS     GlanceShowImageListResponseBodyPlatform
}

func GetGlanceShowImageListResponseBodyPlatformEnum() GlanceShowImageListResponseBodyPlatformEnum {
	return GlanceShowImageListResponseBodyPlatformEnum{
		WINDOWS: GlanceShowImageListResponseBodyPlatform{
			value: "Windows",
		},
		UBUNTU: GlanceShowImageListResponseBodyPlatform{
			value: "Ubuntu",
		},
		RED_HAT: GlanceShowImageListResponseBodyPlatform{
			value: "RedHat",
		},
		SUSE: GlanceShowImageListResponseBodyPlatform{
			value: "SUSE",
		},
		CENT_OS: GlanceShowImageListResponseBodyPlatform{
			value: "CentOS",
		},
		DEBIAN: GlanceShowImageListResponseBodyPlatform{
			value: "Debian",
		},
		OPEN_SUSE: GlanceShowImageListResponseBodyPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: GlanceShowImageListResponseBodyPlatform{
			value: "OracleLinux",
		},
		FEDORA: GlanceShowImageListResponseBodyPlatform{
			value: "Fedora",
		},
		OTHER: GlanceShowImageListResponseBodyPlatform{
			value: "Other",
		},
		CORE_OS: GlanceShowImageListResponseBodyPlatform{
			value: "CoreOS",
		},
		EULER_OS: GlanceShowImageListResponseBodyPlatform{
			value: "EulerOS",
		},
	}
}

func (c GlanceShowImageListResponseBodyPlatform) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyPlatform) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyDiskFormat struct {
	value string
}

type GlanceShowImageListResponseBodyDiskFormatEnum struct {
	VHD   GlanceShowImageListResponseBodyDiskFormat
	ZVHD  GlanceShowImageListResponseBodyDiskFormat
	RAW   GlanceShowImageListResponseBodyDiskFormat
	QCOW2 GlanceShowImageListResponseBodyDiskFormat
	ZVHD2 GlanceShowImageListResponseBodyDiskFormat
}

func GetGlanceShowImageListResponseBodyDiskFormatEnum() GlanceShowImageListResponseBodyDiskFormatEnum {
	return GlanceShowImageListResponseBodyDiskFormatEnum{
		VHD: GlanceShowImageListResponseBodyDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceShowImageListResponseBodyDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceShowImageListResponseBodyDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceShowImageListResponseBodyDiskFormat{
			value: "qcow2",
		},
		ZVHD2: GlanceShowImageListResponseBodyDiskFormat{
			value: "zvhd2",
		},
	}
}

func (c GlanceShowImageListResponseBodyDiskFormat) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyStatus struct {
	value string
}

type GlanceShowImageListResponseBodyStatusEnum struct {
	QUEUED  GlanceShowImageListResponseBodyStatus
	SAVING  GlanceShowImageListResponseBodyStatus
	DELETED GlanceShowImageListResponseBodyStatus
	KILLED  GlanceShowImageListResponseBodyStatus
	ACTIVE  GlanceShowImageListResponseBodyStatus
}

func GetGlanceShowImageListResponseBodyStatusEnum() GlanceShowImageListResponseBodyStatusEnum {
	return GlanceShowImageListResponseBodyStatusEnum{
		QUEUED: GlanceShowImageListResponseBodyStatus{
			value: "queued",
		},
		SAVING: GlanceShowImageListResponseBodyStatus{
			value: "saving",
		},
		DELETED: GlanceShowImageListResponseBodyStatus{
			value: "deleted",
		},
		KILLED: GlanceShowImageListResponseBodyStatus{
			value: "killed",
		},
		ACTIVE: GlanceShowImageListResponseBodyStatus{
			value: "active",
		},
	}
}

func (c GlanceShowImageListResponseBodyStatus) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyStatus) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyVirtualEnvType struct {
	value string
}

type GlanceShowImageListResponseBodyVirtualEnvTypeEnum struct {
	FUSION_COMPUTE GlanceShowImageListResponseBodyVirtualEnvType
	IRONIC         GlanceShowImageListResponseBodyVirtualEnvType
	DATA_IMAGE     GlanceShowImageListResponseBodyVirtualEnvType
}

func GetGlanceShowImageListResponseBodyVirtualEnvTypeEnum() GlanceShowImageListResponseBodyVirtualEnvTypeEnum {
	return GlanceShowImageListResponseBodyVirtualEnvTypeEnum{
		FUSION_COMPUTE: GlanceShowImageListResponseBodyVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: GlanceShowImageListResponseBodyVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: GlanceShowImageListResponseBodyVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c GlanceShowImageListResponseBodyVirtualEnvType) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyVisibility struct {
	value string
}

type GlanceShowImageListResponseBodyVisibilityEnum struct {
	PRIVATE GlanceShowImageListResponseBodyVisibility
	PUBLIC  GlanceShowImageListResponseBodyVisibility
	SHARED  GlanceShowImageListResponseBodyVisibility
}

func GetGlanceShowImageListResponseBodyVisibilityEnum() GlanceShowImageListResponseBodyVisibilityEnum {
	return GlanceShowImageListResponseBodyVisibilityEnum{
		PRIVATE: GlanceShowImageListResponseBodyVisibility{
			value: "private",
		},
		PUBLIC: GlanceShowImageListResponseBodyVisibility{
			value: "public",
		},
		SHARED: GlanceShowImageListResponseBodyVisibility{
			value: "shared",
		},
	}
}

func (c GlanceShowImageListResponseBodyVisibility) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyVisibility) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodySupportFcInject struct {
	value string
}

type GlanceShowImageListResponseBodySupportFcInjectEnum struct {
	TRUE  GlanceShowImageListResponseBodySupportFcInject
	FALSE GlanceShowImageListResponseBodySupportFcInject
}

func GetGlanceShowImageListResponseBodySupportFcInjectEnum() GlanceShowImageListResponseBodySupportFcInjectEnum {
	return GlanceShowImageListResponseBodySupportFcInjectEnum{
		TRUE: GlanceShowImageListResponseBodySupportFcInject{
			value: "true",
		},
		FALSE: GlanceShowImageListResponseBodySupportFcInject{
			value: "false",
		},
	}
}

func (c GlanceShowImageListResponseBodySupportFcInject) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodySupportFcInject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodySupportFcInject) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyHwFirmwareType struct {
	value string
}

type GlanceShowImageListResponseBodyHwFirmwareTypeEnum struct {
	BIOS GlanceShowImageListResponseBodyHwFirmwareType
	UEFI GlanceShowImageListResponseBodyHwFirmwareType
}

func GetGlanceShowImageListResponseBodyHwFirmwareTypeEnum() GlanceShowImageListResponseBodyHwFirmwareTypeEnum {
	return GlanceShowImageListResponseBodyHwFirmwareTypeEnum{
		BIOS: GlanceShowImageListResponseBodyHwFirmwareType{
			value: "bios",
		},
		UEFI: GlanceShowImageListResponseBodyHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c GlanceShowImageListResponseBodyHwFirmwareType) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyHwFirmwareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyHwFirmwareType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodySupportArm struct {
	value string
}

type GlanceShowImageListResponseBodySupportArmEnum struct {
	TRUE  GlanceShowImageListResponseBodySupportArm
	FALSE GlanceShowImageListResponseBodySupportArm
}

func GetGlanceShowImageListResponseBodySupportArmEnum() GlanceShowImageListResponseBodySupportArmEnum {
	return GlanceShowImageListResponseBodySupportArmEnum{
		TRUE: GlanceShowImageListResponseBodySupportArm{
			value: "true",
		},
		FALSE: GlanceShowImageListResponseBodySupportArm{
			value: "false",
		},
	}
}

func (c GlanceShowImageListResponseBodySupportArm) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodySupportArm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodySupportArm) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageListResponseBodyIsOffshelved struct {
	value string
}

type GlanceShowImageListResponseBodyIsOffshelvedEnum struct {
	TRUE  GlanceShowImageListResponseBodyIsOffshelved
	FALSE GlanceShowImageListResponseBodyIsOffshelved
}

func GetGlanceShowImageListResponseBodyIsOffshelvedEnum() GlanceShowImageListResponseBodyIsOffshelvedEnum {
	return GlanceShowImageListResponseBodyIsOffshelvedEnum{
		TRUE: GlanceShowImageListResponseBodyIsOffshelved{
			value: "true",
		},
		FALSE: GlanceShowImageListResponseBodyIsOffshelved{
			value: "false",
		},
	}
}

func (c GlanceShowImageListResponseBodyIsOffshelved) Value() string {
	return c.value
}

func (c GlanceShowImageListResponseBodyIsOffshelved) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageListResponseBodyIsOffshelved) UnmarshalJSON(b []byte) error {
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
