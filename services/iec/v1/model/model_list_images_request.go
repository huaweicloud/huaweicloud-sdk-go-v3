package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListImagesRequest Request Object
type ListImagesRequest struct {

	// 镜像类型，目前支持以下类型：  - 公共镜像：gold  - 私有镜像：private
	Imagetype *string `json:"__imagetype,omitempty"`

	// 镜像是否是受保护，取值为true/false，一般查询公共镜像时候取值为true，查询私有镜像可以不指定。
	Protected *ListImagesRequestProtected `json:"protected,omitempty"`

	// 镜像ID，精确匹配。
	Id *string `json:"id,omitempty"`

	// 是否被其他租户可见，取值如下：  - public：公共镜像  - private：私有镜像
	Visibility *string `json:"visibility,omitempty"`

	// 镜像状态。取值如下：  - saving：表示镜像正在上传文件到后端存储  - deleted：表示镜像已经删除  - killed：表示镜像上传错误  - active：表示镜像可以正常使用
	Status *string `json:"status,omitempty"`

	// 镜像名称，匹配规则为精确匹配。
	Name *string `json:"name,omitempty"`

	// 镜像系统类型，取值如下：  - Linux - Windows - Other
	OsType *string `json:"__os_type,omitempty"`

	// 镜像使用环境类型。  目前仅支持系统盘镜像，取值为：FusionCompute
	VirtualEnvType *ListImagesRequestVirtualEnvType `json:"virtual_env_type,omitempty"`

	// 镜像是否可用，取值为true/false。 > 查询公共镜像时，该参数无效。
	Isregistered *ListImagesRequestIsregistered `json:"__isregistered,omitempty"`

	// 用于分页，表示查询几条镜像记录，取值为正整数，最大（默认）取值为500
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页，表示查询偏移量，表示从整个列表查询结果中的该位置向后进行查询，并非页数偏移，默认取值为0，表示查询第一页。
	Offset *int32 `json:"offset,omitempty"`

	// 用于排序，表示按照哪个字段排序，取值为镜像属性name、status、disk_format、created_at，默认取值为created_at。
	SortKey *ListImagesRequestSortKey `json:"sort_key,omitempty"`

	// 用于排序，表示升序还是降序，取值为asc和desc，与sort_key一起组合使用，默认为降序desc。
	SortDir *ListImagesRequestSortDir `json:"sort_dir,omitempty"`

	// 如果镜像支持KVM，取值为true，否则无该属性。
	SupportKvm *string `json:"__support_kvm,omitempty"`

	// 如果镜像是支持KVM虚拟化平台下的GPU类型，取值为“V100_vGPU”或者“RTX5000”，否则无该属性。
	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty"`

	// 如果镜像支持AI加速，取值为true，否则无该属性。
	SupportKvmAscend310 *string `json:"__support_kvm_ascend_310,omitempty"`

	// 如果镜像支持计算增强，取值为true，否则无该属性。
	SupportKvmHi1822Hiovs *string `json:"__support_kvm_hi1822_hiovs,omitempty"`

	// 如果镜像为ARM架构类型，取值为true，否则无该属性。
	SupportArm *string `json:"__support_arm,omitempty"`

	// 如果镜像支持GPU T4，取值为true，否则无该属性。
	SupportGpuT4 *string `json:"__support_gpu_t4,omitempty"`
}

func (o ListImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesRequest struct{}"
	}

	return strings.Join([]string{"ListImagesRequest", string(data)}, " ")
}

type ListImagesRequestProtected struct {
	value string
}

type ListImagesRequestProtectedEnum struct {
	TRUE  ListImagesRequestProtected
	FALSE ListImagesRequestProtected
}

func GetListImagesRequestProtectedEnum() ListImagesRequestProtectedEnum {
	return ListImagesRequestProtectedEnum{
		TRUE: ListImagesRequestProtected{
			value: "true",
		},
		FALSE: ListImagesRequestProtected{
			value: "false",
		},
	}
}

func (c ListImagesRequestProtected) Value() string {
	return c.value
}

func (c ListImagesRequestProtected) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestProtected) UnmarshalJSON(b []byte) error {
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
}

func GetListImagesRequestVirtualEnvTypeEnum() ListImagesRequestVirtualEnvTypeEnum {
	return ListImagesRequestVirtualEnvTypeEnum{
		FUSION_COMPUTE: ListImagesRequestVirtualEnvType{
			value: "FusionCompute",
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

type ListImagesRequestSortKey struct {
	value string
}

type ListImagesRequestSortKeyEnum struct {
	NAME        ListImagesRequestSortKey
	STATUS      ListImagesRequestSortKey
	DISK_FORMAT ListImagesRequestSortKey
	CREATED_AT  ListImagesRequestSortKey
}

func GetListImagesRequestSortKeyEnum() ListImagesRequestSortKeyEnum {
	return ListImagesRequestSortKeyEnum{
		NAME: ListImagesRequestSortKey{
			value: "name",
		},
		STATUS: ListImagesRequestSortKey{
			value: "status",
		},
		DISK_FORMAT: ListImagesRequestSortKey{
			value: "disk_format",
		},
		CREATED_AT: ListImagesRequestSortKey{
			value: "created_at",
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
