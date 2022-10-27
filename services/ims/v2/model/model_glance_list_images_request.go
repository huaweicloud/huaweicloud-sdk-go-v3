package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type GlanceListImagesRequest struct {

	// 镜像类型，目前支持以下类型：公共镜像：gold私有镜像：private共享镜像：shared 市场镜像：market
	Imagetype *GlanceListImagesRequestImagetype `json:"__imagetype,omitempty"`

	// 镜像是否是受保护，取值为true/false。一般查询公共镜像时候取值为true，查询私有镜像可以不指定。
	Isregistered *bool `json:"__isregistered,omitempty"`

	// 操作系统位数，一般取值为32或者64
	OsBit *GlanceListImagesRequestOsBit `json:"__os_bit,omitempty"`

	// 镜像系统类型，取值为Linux，Windows，Other
	OsType *GlanceListImagesRequestOsType `json:"__os_type,omitempty"`

	// 镜像平台分类，取值为Windows，Ubuntu，RedHat，SUSE，CentOS，Debian，OpenSUSE, Oracle Linux，Fedora，Other，CoreOS和EulerOS
	Platform *GlanceListImagesRequestPlatform `json:"__platform,omitempty"`

	// 表示该镜像支持密集存储。如果镜像支持密集存储性能，则值为true，否则无需增加该属性
	SupportDiskintensive *string `json:"__support_diskintensive,omitempty"`

	// 表示该镜像支持高计算性能。如果镜像支持高计算性能，则值为true，否则无需增加该属性
	SupportHighperformance *string `json:"__support_highperformance,omitempty"`

	// 如果镜像支持KVM，取值为true，否则无需增加该属性
	SupportKvm *string `json:"__support_kvm,omitempty"`

	// 表示该镜像是支持KVM虚拟化平台下的GPU类型,如果不支持KVM虚拟机下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存
	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty"`

	// 如果镜像支持KVM虚拟化下Infiniband网卡类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”属性不共存。
	SupportKvmInfiniband *string `json:"__support_kvm_infiniband,omitempty"`

	// 表示该镜像支持超大内存。如果镜像支持超大内存，取值为true，否则无需增加该属性
	SupportLargememory *string `json:"__support_largememory,omitempty"`

	// 如果镜像支持XEN，取值为true，否则无需增加该属性
	SupportXen *string `json:"__support_xen,omitempty"`

	// 表示该镜像是支持XEN虚拟化平台下的GPU优化类型,如果不支持XEN虚拟化下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存
	SupportXenGpuType *string `json:"__support_xen_gpu_type,omitempty"`

	// 如果镜像支持XEN虚拟化下HANA类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存
	SupportXenHana *string `json:"__support_xen_hana,omitempty"`

	// 容器类型
	ContainerFormat *string `json:"container_format,omitempty"`

	// 镜像格式，目前支持vhd，zvhd、raw，qcow2。默认值是vhd
	DiskFormat *GlanceListImagesRequestDiskFormat `json:"disk_format,omitempty"`

	// 镜像ID
	Id *string `json:"id,omitempty"`

	// 用于分页，表示查询几条镜像记录，取值为整数，默认返回25条镜像记录
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页，表示从哪个镜像开始查询，取值为镜像ID。
	Marker *string `json:"marker,omitempty"`

	// 成员状态。目前取值有accepted、rejected、pending。accepted表示已经接受共享的镜像，rejected表示已经拒绝了其他用户共享的镜像，pending表示需要确认的其他用户的共享镜像。需要在查询时，设置“visibility”参数为“shared”
	MemberStatus *string `json:"member_status,omitempty"`

	// 镜像运行需要的最小磁盘，单位为GB 。取值为40～1024GB。取值为1～1024GB。取值为40～255GB
	MinDisk *int32 `json:"min_disk,omitempty"`

	// 镜像运行需要的最小内存，单位为MB。参数取值依据弹性云服务器的规格限制，一般设置为0。
	MinRam *int32 `json:"min_ram,omitempty"`

	// 镜像名称
	Name *string `json:"name,omitempty"`

	// 镜像属于哪个租户
	Owner *string `json:"owner,omitempty"`

	// 镜像是否是受保护，查询公共镜像时候取值为True，查询私有镜像可以不指定。
	Protected *bool `json:"protected,omitempty"`

	// 用于排序，表示升序还是降序，取值为asc和desc。与sort_key一起组合使用，默认为降序desc
	SortDir *string `json:"sort_dir,omitempty"`

	// 用于排序，表示按照哪个字段排序。取值为镜像属性name，container_format，disk_format，status，id，size字段，默认为创建时间。
	SortKey *string `json:"sort_key,omitempty"`

	// 镜像状态。取值如下：queued：表示镜像元数据已经创建成功，等待上传镜像文件。saving：表示镜像正在上传文件到后端存储。deleted：表示镜像已经删除。killed：表示镜像上传错误。active：表示镜像可以正常使用
	Status *GlanceListImagesRequestStatus `json:"status,omitempty"`

	// 标签，用户为镜像增加自定义标签后可以通过该参数过滤查询
	Tag *string `json:"tag,omitempty"`

	// 是否被其他租户可见，取值如下： public：公共镜像 private：私有镜像 shared：共享镜像
	Visibility *GlanceListImagesRequestVisibility `json:"visibility,omitempty"`

	// 镜像创建时间。支持按照时间点过滤查询，取值格式为“ 操作符:UTC时间”。 其中操作符支持如下几种： gt：大于 gte：大于等于 lt：小于 lte：小于等于 eq：等于 neq：不等于 时间格式支持：yyyy-MM-ddThh:mm:ssZ或者yyyy-MM-dd hh:mm:ss 例如，查询创建时间在2018-10-28 10:00:00之前的镜像，可以通过如下条件过滤： created_at=gt:2018-10-28T10:00:00Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 镜像修改时间。支持按照时间点过滤查询，取值格式为 “ 操作符:UTC时间”。 其中操作符支持如下几种： gt：大于 gte：大于等于 lt：小于 lte：小于等于 eq：等于 neq：不等于 时间格式支持：yyyy-MM-ddThh:mm:ssZ或者yyyy-MM-dd hh:mm:ss 例如，查询修改时间在2018-10-28 10:00:00之前的镜像，可以通过如下条件过滤： updated_at=gt:2018-10-28T10:00:00Z
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o GlanceListImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImagesRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImagesRequest", string(data)}, " ")
}

type GlanceListImagesRequestImagetype struct {
	value string
}

type GlanceListImagesRequestImagetypeEnum struct {
	GOLD    GlanceListImagesRequestImagetype
	PRIVATE GlanceListImagesRequestImagetype
	SHARED  GlanceListImagesRequestImagetype
	MARKET  GlanceListImagesRequestImagetype
}

func GetGlanceListImagesRequestImagetypeEnum() GlanceListImagesRequestImagetypeEnum {
	return GlanceListImagesRequestImagetypeEnum{
		GOLD: GlanceListImagesRequestImagetype{
			value: "gold",
		},
		PRIVATE: GlanceListImagesRequestImagetype{
			value: "private",
		},
		SHARED: GlanceListImagesRequestImagetype{
			value: "shared",
		},
		MARKET: GlanceListImagesRequestImagetype{
			value: "market",
		},
	}
}

func (c GlanceListImagesRequestImagetype) Value() string {
	return c.value
}

func (c GlanceListImagesRequestImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestImagetype) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestOsBit struct {
	value string
}

type GlanceListImagesRequestOsBitEnum struct {
	E_32 GlanceListImagesRequestOsBit
	E_64 GlanceListImagesRequestOsBit
}

func GetGlanceListImagesRequestOsBitEnum() GlanceListImagesRequestOsBitEnum {
	return GlanceListImagesRequestOsBitEnum{
		E_32: GlanceListImagesRequestOsBit{
			value: "32",
		},
		E_64: GlanceListImagesRequestOsBit{
			value: "64",
		},
	}
}

func (c GlanceListImagesRequestOsBit) Value() string {
	return c.value
}

func (c GlanceListImagesRequestOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestOsType struct {
	value string
}

type GlanceListImagesRequestOsTypeEnum struct {
	LINUX   GlanceListImagesRequestOsType
	WINDOWS GlanceListImagesRequestOsType
	OTHER   GlanceListImagesRequestOsType
}

func GetGlanceListImagesRequestOsTypeEnum() GlanceListImagesRequestOsTypeEnum {
	return GlanceListImagesRequestOsTypeEnum{
		LINUX: GlanceListImagesRequestOsType{
			value: "Linux",
		},
		WINDOWS: GlanceListImagesRequestOsType{
			value: "Windows",
		},
		OTHER: GlanceListImagesRequestOsType{
			value: "Other",
		},
	}
}

func (c GlanceListImagesRequestOsType) Value() string {
	return c.value
}

func (c GlanceListImagesRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestOsType) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestPlatform struct {
	value string
}

type GlanceListImagesRequestPlatformEnum struct {
	WINDOWS      GlanceListImagesRequestPlatform
	UBUNTU       GlanceListImagesRequestPlatform
	RED_HAT      GlanceListImagesRequestPlatform
	SUSE         GlanceListImagesRequestPlatform
	CENT_OS      GlanceListImagesRequestPlatform
	DEBIAN       GlanceListImagesRequestPlatform
	OPEN_SUSE    GlanceListImagesRequestPlatform
	ORACLE_LINUX GlanceListImagesRequestPlatform
	FEDORA       GlanceListImagesRequestPlatform
	OTHER        GlanceListImagesRequestPlatform
	CORE_OS      GlanceListImagesRequestPlatform
	EULER_OS     GlanceListImagesRequestPlatform
}

func GetGlanceListImagesRequestPlatformEnum() GlanceListImagesRequestPlatformEnum {
	return GlanceListImagesRequestPlatformEnum{
		WINDOWS: GlanceListImagesRequestPlatform{
			value: "Windows",
		},
		UBUNTU: GlanceListImagesRequestPlatform{
			value: "Ubuntu",
		},
		RED_HAT: GlanceListImagesRequestPlatform{
			value: "RedHat",
		},
		SUSE: GlanceListImagesRequestPlatform{
			value: "SUSE",
		},
		CENT_OS: GlanceListImagesRequestPlatform{
			value: "CentOS",
		},
		DEBIAN: GlanceListImagesRequestPlatform{
			value: "Debian",
		},
		OPEN_SUSE: GlanceListImagesRequestPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: GlanceListImagesRequestPlatform{
			value: "Oracle Linux",
		},
		FEDORA: GlanceListImagesRequestPlatform{
			value: "Fedora",
		},
		OTHER: GlanceListImagesRequestPlatform{
			value: "Other",
		},
		CORE_OS: GlanceListImagesRequestPlatform{
			value: "CoreOS",
		},
		EULER_OS: GlanceListImagesRequestPlatform{
			value: "EulerOS",
		},
	}
}

func (c GlanceListImagesRequestPlatform) Value() string {
	return c.value
}

func (c GlanceListImagesRequestPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestPlatform) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestDiskFormat struct {
	value string
}

type GlanceListImagesRequestDiskFormatEnum struct {
	VHD   GlanceListImagesRequestDiskFormat
	ZVHD  GlanceListImagesRequestDiskFormat
	RAW   GlanceListImagesRequestDiskFormat
	QCOW2 GlanceListImagesRequestDiskFormat
}

func GetGlanceListImagesRequestDiskFormatEnum() GlanceListImagesRequestDiskFormatEnum {
	return GlanceListImagesRequestDiskFormatEnum{
		VHD: GlanceListImagesRequestDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceListImagesRequestDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceListImagesRequestDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceListImagesRequestDiskFormat{
			value: "qcow2",
		},
	}
}

func (c GlanceListImagesRequestDiskFormat) Value() string {
	return c.value
}

func (c GlanceListImagesRequestDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestStatus struct {
	value string
}

type GlanceListImagesRequestStatusEnum struct {
	QUEUED  GlanceListImagesRequestStatus
	SAVING  GlanceListImagesRequestStatus
	DELETED GlanceListImagesRequestStatus
	KILLED  GlanceListImagesRequestStatus
	ACTIVE  GlanceListImagesRequestStatus
}

func GetGlanceListImagesRequestStatusEnum() GlanceListImagesRequestStatusEnum {
	return GlanceListImagesRequestStatusEnum{
		QUEUED: GlanceListImagesRequestStatus{
			value: "queued",
		},
		SAVING: GlanceListImagesRequestStatus{
			value: "saving",
		},
		DELETED: GlanceListImagesRequestStatus{
			value: "deleted",
		},
		KILLED: GlanceListImagesRequestStatus{
			value: "killed",
		},
		ACTIVE: GlanceListImagesRequestStatus{
			value: "active",
		},
	}
}

func (c GlanceListImagesRequestStatus) Value() string {
	return c.value
}

func (c GlanceListImagesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestStatus) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestVisibility struct {
	value string
}

type GlanceListImagesRequestVisibilityEnum struct {
	PUBLIC  GlanceListImagesRequestVisibility
	PRIVATE GlanceListImagesRequestVisibility
	SHARED  GlanceListImagesRequestVisibility
}

func GetGlanceListImagesRequestVisibilityEnum() GlanceListImagesRequestVisibilityEnum {
	return GlanceListImagesRequestVisibilityEnum{
		PUBLIC: GlanceListImagesRequestVisibility{
			value: "public",
		},
		PRIVATE: GlanceListImagesRequestVisibility{
			value: "private",
		},
		SHARED: GlanceListImagesRequestVisibility{
			value: "shared",
		},
	}
}

func (c GlanceListImagesRequestVisibility) Value() string {
	return c.value
}

func (c GlanceListImagesRequestVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestVisibility) UnmarshalJSON(b []byte) error {
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
