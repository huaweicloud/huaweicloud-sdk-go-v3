package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListCloudImagesRequest struct {

	// 华为云区域ID
	RegionId string `json:"region_id"`

	// 镜像类型，目前支持以下类型：  - 公共镜像：gold  - 私有镜像：private
	Imagetype *string `json:"__imagetype,omitempty"`

	// 镜像是否可用，取值为true/false。 > 查询公共镜像时，该参数无效。
	Isregistered *ListCloudImagesRequestIsregistered `json:"__isregistered,omitempty"`

	// 镜像系统类型，取值如下：  - Linux - Windows - Other
	OsType *string `json:"__os_type,omitempty"`

	// 表示该镜像支持密集存储。如果镜像支持密集存储性能，则值为true，否则无需增加该属性。
	SupportDiskintensive *string `json:"__support_diskintensive,omitempty"`

	// 表示该镜像支持高计算性能。如果镜像支持高计算性能，则值为true，否则无需增加该属性。
	SupportHighperformance *string `json:"__support_highperformance,omitempty"`

	// 如果镜像支持KVM，取值为true，否则无该属性。
	SupportKvm *string `json:"__support_kvm,omitempty"`

	// 如果镜像是支持KVM虚拟化平台下的GPU类型，取值为“V100_vGPU”或者“RTX5000”，否则无该属性。
	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty"`

	// 如果镜像支持KVM虚拟化下Infiniband网卡类型，取值为true。否则，无需添加该属性。  该属性与“__support_xen”属性不共存。
	SupportKvmInfiniband *string `json:"__support_kvm_infiniband,omitempty"`

	// 表示该镜像支持超大内存。如果镜像支持超大内存，取值为true，否则无需增加该属性。
	SupportLargememory *string `json:"__support_largememory,omitempty"`

	// 如果镜像支持XEN，取值为true，否则无需增加该属性。
	SupportXen *string `json:"__support_xen,omitempty"`

	// 表示该镜像是支持XEN虚拟化平台下的GPU优化类型。如果不支持XEN虚拟化下GPU类型，无需添加该属性。该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenGpuType *string `json:"__support_xen_gpu_type,omitempty"`

	// 如果镜像支持XEN虚拟化下HANA类型，取值为true。否则，无需添加该属性。  该属性与“__support_xen”和“__support_kvm”属性不共存。
	SupportXenHana *string `json:"__support_xen_hana,omitempty"`

	// 镜像ID，精确匹配。
	Id *string `json:"id,omitempty"`

	// 用于分页，表示查询几条镜像记录，取值为正整数，最大（默认）取值为500
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页，表示从哪个镜像开始查询，取值为镜像ID。
	Marker *string `json:"marker,omitempty"`

	// 镜像名称，匹配规则为精确匹配。
	Name *string `json:"name,omitempty"`

	// 镜像属于哪个租户。
	Owner *string `json:"owner,omitempty"`

	// 镜像是否是受保护，取值为true/false，一般查询公共镜像时候取值为true，查询私有镜像可以不指定。
	Protected *ListCloudImagesRequestProtected `json:"protected,omitempty"`

	// 用于排序，表示升序还是降序，取值为asc和desc，与sort_key一起组合使用，默认为降序desc。
	SortDir *ListCloudImagesRequestSortDir `json:"sort_dir,omitempty"`

	// 用于排序，表示按照哪个字段排序，取值为镜像属性name、status、disk_format、created_at，默认取值为created_at。
	SortKey *ListCloudImagesRequestSortKey `json:"sort_key,omitempty"`

	// 镜像状态。取值如下：  - saving：表示镜像正在上传文件到后端存储  - deleted：表示镜像已经删除  - killed：表示镜像上传错误  - active：表示镜像可以正常使用
	Status *string `json:"status,omitempty"`

	// 镜像使用环境类型。  目前仅支持系统盘镜像，取值为：FusionCompute
	VirtualEnvType *ListCloudImagesRequestVirtualEnvType `json:"virtual_env_type,omitempty"`

	// 是否被其他租户可见，取值如下：  - public：公共镜像  - private：私有镜像
	Visibility *string `json:"visibility,omitempty"`
}

func (o ListCloudImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudImagesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudImagesRequest", string(data)}, " ")
}

type ListCloudImagesRequestIsregistered struct {
	value string
}

type ListCloudImagesRequestIsregisteredEnum struct {
	TRUE  ListCloudImagesRequestIsregistered
	FALSE ListCloudImagesRequestIsregistered
}

func GetListCloudImagesRequestIsregisteredEnum() ListCloudImagesRequestIsregisteredEnum {
	return ListCloudImagesRequestIsregisteredEnum{
		TRUE: ListCloudImagesRequestIsregistered{
			value: "true",
		},
		FALSE: ListCloudImagesRequestIsregistered{
			value: "false",
		},
	}
}

func (c ListCloudImagesRequestIsregistered) Value() string {
	return c.value
}

func (c ListCloudImagesRequestIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloudImagesRequestIsregistered) UnmarshalJSON(b []byte) error {
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

type ListCloudImagesRequestProtected struct {
	value string
}

type ListCloudImagesRequestProtectedEnum struct {
	TRUE  ListCloudImagesRequestProtected
	FALSE ListCloudImagesRequestProtected
}

func GetListCloudImagesRequestProtectedEnum() ListCloudImagesRequestProtectedEnum {
	return ListCloudImagesRequestProtectedEnum{
		TRUE: ListCloudImagesRequestProtected{
			value: "true",
		},
		FALSE: ListCloudImagesRequestProtected{
			value: "false",
		},
	}
}

func (c ListCloudImagesRequestProtected) Value() string {
	return c.value
}

func (c ListCloudImagesRequestProtected) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloudImagesRequestProtected) UnmarshalJSON(b []byte) error {
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

type ListCloudImagesRequestSortDir struct {
	value string
}

type ListCloudImagesRequestSortDirEnum struct {
	ASC  ListCloudImagesRequestSortDir
	DESC ListCloudImagesRequestSortDir
}

func GetListCloudImagesRequestSortDirEnum() ListCloudImagesRequestSortDirEnum {
	return ListCloudImagesRequestSortDirEnum{
		ASC: ListCloudImagesRequestSortDir{
			value: "asc",
		},
		DESC: ListCloudImagesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListCloudImagesRequestSortDir) Value() string {
	return c.value
}

func (c ListCloudImagesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloudImagesRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListCloudImagesRequestSortKey struct {
	value string
}

type ListCloudImagesRequestSortKeyEnum struct {
	NAME        ListCloudImagesRequestSortKey
	STATUS      ListCloudImagesRequestSortKey
	DISK_FORMAT ListCloudImagesRequestSortKey
	CREATED_AT  ListCloudImagesRequestSortKey
}

func GetListCloudImagesRequestSortKeyEnum() ListCloudImagesRequestSortKeyEnum {
	return ListCloudImagesRequestSortKeyEnum{
		NAME: ListCloudImagesRequestSortKey{
			value: "name",
		},
		STATUS: ListCloudImagesRequestSortKey{
			value: "status",
		},
		DISK_FORMAT: ListCloudImagesRequestSortKey{
			value: "disk_format",
		},
		CREATED_AT: ListCloudImagesRequestSortKey{
			value: "created_at",
		},
	}
}

func (c ListCloudImagesRequestSortKey) Value() string {
	return c.value
}

func (c ListCloudImagesRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloudImagesRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListCloudImagesRequestVirtualEnvType struct {
	value string
}

type ListCloudImagesRequestVirtualEnvTypeEnum struct {
	FUSION_COMPUTE ListCloudImagesRequestVirtualEnvType
}

func GetListCloudImagesRequestVirtualEnvTypeEnum() ListCloudImagesRequestVirtualEnvTypeEnum {
	return ListCloudImagesRequestVirtualEnvTypeEnum{
		FUSION_COMPUTE: ListCloudImagesRequestVirtualEnvType{
			value: "FusionCompute",
		},
	}
}

func (c ListCloudImagesRequestVirtualEnvType) Value() string {
	return c.value
}

func (c ListCloudImagesRequestVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloudImagesRequestVirtualEnvType) UnmarshalJSON(b []byte) error {
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
