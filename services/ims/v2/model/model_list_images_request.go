package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListImagesRequest Request Object
type ListImagesRequest struct {

	// 镜像类型，目前支持以下类型： 公共镜像：gold 私有镜像：private 共享镜像：shared 市场镜像：market
	Imagetype *ListImagesRequestImagetype `json:"__imagetype,omitempty"`

	// 镜像是否可用，取值为true，扩展接口会默认为true，普通用户只能查询取值为true的镜像。
	Isregistered *ListImagesRequestIsregistered `json:"__isregistered,omitempty"`

	// 操作系统位数，一般取值为32或者64。
	OsBit *ListImagesRequestOsBit `json:"__os_bit,omitempty"`

	// 镜像系统类型，取值为Linux，Windows，Other。
	OsType *ListImagesRequestOsType `json:"__os_type,omitempty"`

	// 镜像平台分类
	Platform *ListImagesRequestPlatform `json:"__platform,omitempty"`

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

	// 表示该镜像支持超大内存。如果镜像支持超大内存，取值为true，否则无需增加该属性。
	SupportLargememory *string `json:"__support_largememory,omitempty"`

	// 如果镜像支持XEN，取值为true，否则无需增加该属性。
	SupportXen *string `json:"__support_xen,omitempty"`

	// 表示该镜像是支持XEN虚拟化平台下的GPU优化类型，如果不支持XEN虚拟化下GPU类型，无需添加该属性 。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenGpuType *string `json:"__support_xen_gpu_type,omitempty"`

	// 如果镜像支持XEN虚拟化下HANA类型，取值为true。否则，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenHana *string `json:"__support_xen_hana,omitempty"`

	// 容器类型
	ContainerFormat *string `json:"container_format,omitempty"`

	// 镜像格式，目前支持vhd，zvhd、raw，qcow2,zvhd2。默认值是vhd。
	DiskFormat *ListImagesRequestDiskFormat `json:"disk_format,omitempty"`

	// 表示查询某个企业项目下的镜像。 取值为0，表示查询属于default企业项目下的镜像。 取值为UUID，表示查询属于该UUID对应的企业项目下的镜像。取值为all_granted_eps，表示查询当前用户所有企业项目下的镜像。 关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像ID
	Id *string `json:"id,omitempty"`

	// 用于分页，表示查询几条镜像记录，取值为整数，默认取值为500。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页，表示从哪个镜像开始查询，取值为镜像ID。
	Marker *string `json:"marker,omitempty"`

	// 成员状态。目前取值有accepted、rejected、pending。accepted表示已经接受共享的镜像，rejected表示已经拒绝了其他用户共享的镜像，pending表示需要确认的其他用户的共享镜像。需要在查询时设置“visibility”参数为“shared”。
	MemberStatus *ListImagesRequestMemberStatus `json:"member_status,omitempty"`

	// 镜像运行需要的最小磁盘，单位为GB 。取值为40～1024GB。
	MinDisk *int32 `json:"min_disk,omitempty"`

	// 镜像运行需要的最小内存，单位为MB。参数取值依据弹性云服务器的规格限制，一般设置为0。
	MinRam *int32 `json:"min_ram,omitempty"`

	// 镜像名称
	Name *string `json:"name,omitempty"`

	// 镜像属于哪个租户
	Owner *string `json:"owner,omitempty"`

	// 镜像是否是受保护，取值为true/false，一般查询公共镜像时候取值为true，查询私有镜像可以不指定。
	Protected *bool `json:"protected,omitempty"`

	// 用于排序，表示升序还是降序，取值为asc和desc。与sort_key一起组合使用，默认为降序desc。
	SortDir *ListImagesRequestSortDir `json:"sort_dir,omitempty"`

	// 用于排序，表示按照哪个字段排序。取值为镜像属性name，container_format，disk_format，status ，id，size字段，默认为创建时间。
	SortKey *ListImagesRequestSortKey `json:"sort_key,omitempty"`

	// 镜像状态。取值如下： queued：表示镜像元数据已经创建成功，等待上传镜像文件。 saving：表示镜像正在上传文件到后端存储。 deleted：表示镜像已经删除。 killed：表示镜像上传错误。 active：表示镜像可以正常使用。
	Status *ListImagesRequestStatus `json:"status,omitempty"`

	// 标签，用户为镜像增加自定义标签后可以通过该参数过滤查询。
	Tag *string `json:"tag,omitempty"`

	// 镜像使用环境类型：FusionCompute，Ironic，DataImage。如果弹性云服务器镜像，则取值为FusionCompute，如果是数据卷镜像则取值是DataImage，如果是裸金属服务器镜像，则取值是Ironic。
	VirtualEnvType *ListImagesRequestVirtualEnvType `json:"virtual_env_type,omitempty"`

	// 是否被其他租户可见，取值为public、private或shared
	Visibility *ListImagesRequestVisibility `json:"visibility,omitempty"`

	// 请求的发生时间,格式为YYYYMMDDTHHMMSSZ。取值为当前系统的GMT时间。使用AK/SK认证时该字段必选
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 用于通过云服务器规格过滤出可用公共镜像，取值为规格ID。 当前仅支持通过单个规格进行过滤。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 镜像创建时间。支持按照时间点过滤查询，取值格式为“操作符:UTC时间”。 其中操作符支持如下几种： gt：大于 gte：大于等于 lt：小于 lte：小于等于 eq：等于 neq：不等于 时间格式支持：yyyy-MM-ddThh:mm:ssZ或者yyyy-MM-dd hh:mm:ss 例如，查询创建时间在2018-10-28 10:00:00之前的镜像，可以通过如下条件过滤： created_at=gt:2018-10-28T10:00:00Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 镜像修改时间。支持按照时间点过滤查询，取值格式为“ 操作符:UTC时间”。 其中操作符支持如下几种： gt：大于 gte：大于等于 lt：小于 lte：小于等于 eq：等于 neq：不等于 时间格式支持：yyyy-MM-ddThh:mm:ssZ或者yyyy-MM-dd hh:mm:ss 例如，查询修改时间在2018-10-28 10:00:00之前的镜像，可以通过如下条件过滤： updated_at=gt:2018-10-28T10:00:00Z
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 镜像架构类型。取值包括： x86 arm
	Architecture *ListImagesRequestArchitecture `json:"architecture,omitempty"`
}

func (o ListImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesRequest struct{}"
	}

	return strings.Join([]string{"ListImagesRequest", string(data)}, " ")
}

type ListImagesRequestImagetype struct {
	value string
}

type ListImagesRequestImagetypeEnum struct {
	GOLD    ListImagesRequestImagetype
	PRIVATE ListImagesRequestImagetype
	SHARED  ListImagesRequestImagetype
	MARKET  ListImagesRequestImagetype
}

func GetListImagesRequestImagetypeEnum() ListImagesRequestImagetypeEnum {
	return ListImagesRequestImagetypeEnum{
		GOLD: ListImagesRequestImagetype{
			value: "gold",
		},
		PRIVATE: ListImagesRequestImagetype{
			value: "private",
		},
		SHARED: ListImagesRequestImagetype{
			value: "shared",
		},
		MARKET: ListImagesRequestImagetype{
			value: "market",
		},
	}
}

func (c ListImagesRequestImagetype) Value() string {
	return c.value
}

func (c ListImagesRequestImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestImagetype) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestIsregistered struct {
	value string
}

type ListImagesRequestIsregisteredEnum struct {
	TRUE  ListImagesRequestIsregistered
	FALSE ListImagesRequestIsregistered
}

func GetListImagesRequestIsregisteredEnum() ListImagesRequestIsregisteredEnum {
	return ListImagesRequestIsregisteredEnum{
		TRUE: ListImagesRequestIsregistered{
			value: "true",
		},
		FALSE: ListImagesRequestIsregistered{
			value: "false",
		},
	}
}

func (c ListImagesRequestIsregistered) Value() string {
	return c.value
}

func (c ListImagesRequestIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestIsregistered) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestOsBit struct {
	value string
}

type ListImagesRequestOsBitEnum struct {
	E_32 ListImagesRequestOsBit
	E_64 ListImagesRequestOsBit
}

func GetListImagesRequestOsBitEnum() ListImagesRequestOsBitEnum {
	return ListImagesRequestOsBitEnum{
		E_32: ListImagesRequestOsBit{
			value: "32",
		},
		E_64: ListImagesRequestOsBit{
			value: "64",
		},
	}
}

func (c ListImagesRequestOsBit) Value() string {
	return c.value
}

func (c ListImagesRequestOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestOsBit) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestOsType struct {
	value string
}

type ListImagesRequestOsTypeEnum struct {
	LINUX   ListImagesRequestOsType
	WINDOWS ListImagesRequestOsType
	OTHER   ListImagesRequestOsType
}

func GetListImagesRequestOsTypeEnum() ListImagesRequestOsTypeEnum {
	return ListImagesRequestOsTypeEnum{
		LINUX: ListImagesRequestOsType{
			value: "Linux",
		},
		WINDOWS: ListImagesRequestOsType{
			value: "Windows",
		},
		OTHER: ListImagesRequestOsType{
			value: "Other",
		},
	}
}

func (c ListImagesRequestOsType) Value() string {
	return c.value
}

func (c ListImagesRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestOsType) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestPlatform struct {
	value string
}

type ListImagesRequestPlatformEnum struct {
	WINDOWS               ListImagesRequestPlatform
	UBUNTU                ListImagesRequestPlatform
	RED_HAT               ListImagesRequestPlatform
	SUSE                  ListImagesRequestPlatform
	CENT_OS               ListImagesRequestPlatform
	DEBIAN                ListImagesRequestPlatform
	OPEN_SUSE             ListImagesRequestPlatform
	ORACLE_LINUX          ListImagesRequestPlatform
	FEDORA                ListImagesRequestPlatform
	OTHER                 ListImagesRequestPlatform
	CORE_OS               ListImagesRequestPlatform
	EULER_OS              ListImagesRequestPlatform
	HUAWEI_CLOUD_EULER_OS ListImagesRequestPlatform
}

func GetListImagesRequestPlatformEnum() ListImagesRequestPlatformEnum {
	return ListImagesRequestPlatformEnum{
		WINDOWS: ListImagesRequestPlatform{
			value: "Windows",
		},
		UBUNTU: ListImagesRequestPlatform{
			value: "Ubuntu",
		},
		RED_HAT: ListImagesRequestPlatform{
			value: "RedHat",
		},
		SUSE: ListImagesRequestPlatform{
			value: "SUSE",
		},
		CENT_OS: ListImagesRequestPlatform{
			value: "CentOS",
		},
		DEBIAN: ListImagesRequestPlatform{
			value: "Debian",
		},
		OPEN_SUSE: ListImagesRequestPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: ListImagesRequestPlatform{
			value: "Oracle Linux",
		},
		FEDORA: ListImagesRequestPlatform{
			value: "Fedora",
		},
		OTHER: ListImagesRequestPlatform{
			value: "Other",
		},
		CORE_OS: ListImagesRequestPlatform{
			value: "CoreOS",
		},
		EULER_OS: ListImagesRequestPlatform{
			value: "EulerOS",
		},
		HUAWEI_CLOUD_EULER_OS: ListImagesRequestPlatform{
			value: "Huawei Cloud EulerOS",
		},
	}
}

func (c ListImagesRequestPlatform) Value() string {
	return c.value
}

func (c ListImagesRequestPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestPlatform) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestDiskFormat struct {
	value string
}

type ListImagesRequestDiskFormatEnum struct {
	VHD   ListImagesRequestDiskFormat
	ZVHD  ListImagesRequestDiskFormat
	RAW   ListImagesRequestDiskFormat
	QCOW2 ListImagesRequestDiskFormat
	ZVHD2 ListImagesRequestDiskFormat
}

func GetListImagesRequestDiskFormatEnum() ListImagesRequestDiskFormatEnum {
	return ListImagesRequestDiskFormatEnum{
		VHD: ListImagesRequestDiskFormat{
			value: "vhd",
		},
		ZVHD: ListImagesRequestDiskFormat{
			value: "zvhd",
		},
		RAW: ListImagesRequestDiskFormat{
			value: "raw",
		},
		QCOW2: ListImagesRequestDiskFormat{
			value: "qcow2",
		},
		ZVHD2: ListImagesRequestDiskFormat{
			value: "zvhd2",
		},
	}
}

func (c ListImagesRequestDiskFormat) Value() string {
	return c.value
}

func (c ListImagesRequestDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestDiskFormat) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestMemberStatus struct {
	value string
}

type ListImagesRequestMemberStatusEnum struct {
	ACCEPTED ListImagesRequestMemberStatus
	REJECTED ListImagesRequestMemberStatus
	PENDING  ListImagesRequestMemberStatus
}

func GetListImagesRequestMemberStatusEnum() ListImagesRequestMemberStatusEnum {
	return ListImagesRequestMemberStatusEnum{
		ACCEPTED: ListImagesRequestMemberStatus{
			value: "accepted",
		},
		REJECTED: ListImagesRequestMemberStatus{
			value: "rejected",
		},
		PENDING: ListImagesRequestMemberStatus{
			value: "pending",
		},
	}
}

func (c ListImagesRequestMemberStatus) Value() string {
	return c.value
}

func (c ListImagesRequestMemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestMemberStatus) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestSortDir struct {
	value string
}

type ListImagesRequestSortDirEnum struct {
	ASC  ListImagesRequestSortDir
	DESC ListImagesRequestSortDir
}

func GetListImagesRequestSortDirEnum() ListImagesRequestSortDirEnum {
	return ListImagesRequestSortDirEnum{
		ASC: ListImagesRequestSortDir{
			value: "asc",
		},
		DESC: ListImagesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListImagesRequestSortDir) Value() string {
	return c.value
}

func (c ListImagesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestSortKey struct {
	value string
}

type ListImagesRequestSortKeyEnum struct {
	CREATED_AT       ListImagesRequestSortKey
	NAME             ListImagesRequestSortKey
	CONTAINER_FORMAT ListImagesRequestSortKey
	DISK_FORMAT      ListImagesRequestSortKey
	STATUS           ListImagesRequestSortKey
	ID               ListImagesRequestSortKey
	SIZE             ListImagesRequestSortKey
}

func GetListImagesRequestSortKeyEnum() ListImagesRequestSortKeyEnum {
	return ListImagesRequestSortKeyEnum{
		CREATED_AT: ListImagesRequestSortKey{
			value: "created_at",
		},
		NAME: ListImagesRequestSortKey{
			value: "name",
		},
		CONTAINER_FORMAT: ListImagesRequestSortKey{
			value: "container_format",
		},
		DISK_FORMAT: ListImagesRequestSortKey{
			value: "disk_format",
		},
		STATUS: ListImagesRequestSortKey{
			value: "status ",
		},
		ID: ListImagesRequestSortKey{
			value: "id",
		},
		SIZE: ListImagesRequestSortKey{
			value: "size",
		},
	}
}

func (c ListImagesRequestSortKey) Value() string {
	return c.value
}

func (c ListImagesRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestStatus struct {
	value string
}

type ListImagesRequestStatusEnum struct {
	QUEUED  ListImagesRequestStatus
	SAVING  ListImagesRequestStatus
	DELETED ListImagesRequestStatus
	KILLED  ListImagesRequestStatus
	ACTIVE  ListImagesRequestStatus
}

func GetListImagesRequestStatusEnum() ListImagesRequestStatusEnum {
	return ListImagesRequestStatusEnum{
		QUEUED: ListImagesRequestStatus{
			value: "queued",
		},
		SAVING: ListImagesRequestStatus{
			value: "saving",
		},
		DELETED: ListImagesRequestStatus{
			value: "deleted",
		},
		KILLED: ListImagesRequestStatus{
			value: "killed",
		},
		ACTIVE: ListImagesRequestStatus{
			value: "active",
		},
	}
}

func (c ListImagesRequestStatus) Value() string {
	return c.value
}

func (c ListImagesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestVirtualEnvType struct {
	value string
}

type ListImagesRequestVirtualEnvTypeEnum struct {
	FUSION_COMPUTE ListImagesRequestVirtualEnvType
	IRONIC         ListImagesRequestVirtualEnvType
	DATA_IMAGE     ListImagesRequestVirtualEnvType
}

func GetListImagesRequestVirtualEnvTypeEnum() ListImagesRequestVirtualEnvTypeEnum {
	return ListImagesRequestVirtualEnvTypeEnum{
		FUSION_COMPUTE: ListImagesRequestVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: ListImagesRequestVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: ListImagesRequestVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c ListImagesRequestVirtualEnvType) Value() string {
	return c.value
}

func (c ListImagesRequestVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestVisibility struct {
	value string
}

type ListImagesRequestVisibilityEnum struct {
	PUBLIC  ListImagesRequestVisibility
	PRIVATE ListImagesRequestVisibility
	SHARED  ListImagesRequestVisibility
}

func GetListImagesRequestVisibilityEnum() ListImagesRequestVisibilityEnum {
	return ListImagesRequestVisibilityEnum{
		PUBLIC: ListImagesRequestVisibility{
			value: "public",
		},
		PRIVATE: ListImagesRequestVisibility{
			value: "private",
		},
		SHARED: ListImagesRequestVisibility{
			value: "shared",
		},
	}
}

func (c ListImagesRequestVisibility) Value() string {
	return c.value
}

func (c ListImagesRequestVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestVisibility) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestArchitecture struct {
	value string
}

type ListImagesRequestArchitectureEnum struct {
	X86 ListImagesRequestArchitecture
	ARM ListImagesRequestArchitecture
}

func GetListImagesRequestArchitectureEnum() ListImagesRequestArchitectureEnum {
	return ListImagesRequestArchitectureEnum{
		X86: ListImagesRequestArchitecture{
			value: "x86",
		},
		ARM: ListImagesRequestArchitecture{
			value: "arm",
		},
	}
}

func (c ListImagesRequestArchitecture) Value() string {
	return c.value
}

func (c ListImagesRequestArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestArchitecture) UnmarshalJSON(b []byte) error {
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
