package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type ImageList struct {
	// 镜像ID。

	Id *string `json:"id,omitempty"`
	// 镜像名称。

	Name *string `json:"name,omitempty"`
	// 镜像状态。

	Status *ImageListStatus `json:"status,omitempty"`
	// 镜像格式。

	DiskFormat *ImageListDiskFormat `json:"disk_format,omitempty"`
	// 最小系统盘（单位：GB），取值为40～1024GB。

	MinDisk *int32 `json:"min_disk,omitempty"`
	// 最小内存（单位：MB），默认值为0。

	MinRam *int32 `json:"min_ram,omitempty"`
	// 镜像所属租户ID。

	Owner *string `json:"owner,omitempty"`
	// 是否受保护。

	Protected *bool `json:"protected,omitempty"`
	// 可见性。

	Visibility *string `json:"visibility,omitempty"`
	// 创建时间。

	CreatedAt *string `json:"created_at,omitempty"`
	// 更新时间。

	UpdatedAt *string `json:"updated_at,omitempty"`
	// 镜像链接信息。

	Self *string `json:"self,omitempty"`
	// 是否是删除的镜像，取值为true或者false。

	Deleted *bool `json:"deleted,omitempty"`
	// 镜像使用环境类型。

	VirtualEnvType *string `json:"virtual_env_type,omitempty"`
	// 删除时间，格式为UTC时间。

	DeletedAt *string `json:"deleted_at,omitempty"`
	// 镜像关联的任务ID。

	RelationJobId *string `json:"relation_job_id,omitempty"`
	// 镜像类型。  取值范围： - gold：公有镜像； - private：私有镜像。

	Imagetype *ImageListImagetype `json:"__imagetype,omitempty"`
	// 镜像平台分类。

	Platform *string `json:"__platform,omitempty"`
	// 镜像系统类型。

	OsType *ImageListOsType `json:"__os_type,omitempty"`
	// 镜像的操作系统具体版本。

	OsVersion *string `json:"__os_version,omitempty"`
	// 是否是注册过的镜像。

	Isregistered *bool `json:"__isregistered,omitempty"`
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
	// 镜像启动模式，取值为uefi或bios，不指定时无该属性。

	HwFirmwareType *string `json:"hw_firmware_type,omitempty"`
	// 镜像来源。  - 来源边缘实例：instance:<实例ID> - 来源IMS：ims:<镜像ID>:<region ID>

	DataSource *string `json:"data_source,omitempty"`
	// 如果镜像支持GPU T4类型，取值为true，否则无该属性。

	SupportGpuT4 *string `json:"__support_gpu_t4,omitempty"`
}

func (o ImageList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageList struct{}"
	}

	return strings.Join([]string{"ImageList", string(data)}, " ")
}

type ImageListStatus struct {
	value string
}

type ImageListStatusEnum struct {
	QUEUED  ImageListStatus
	SAVING  ImageListStatus
	DELETED ImageListStatus
	KILLED  ImageListStatus
	ACTIVE  ImageListStatus
}

func GetImageListStatusEnum() ImageListStatusEnum {
	return ImageListStatusEnum{
		QUEUED: ImageListStatus{
			value: "queued",
		},
		SAVING: ImageListStatus{
			value: "saving",
		},
		DELETED: ImageListStatus{
			value: "deleted",
		},
		KILLED: ImageListStatus{
			value: "killed",
		},
		ACTIVE: ImageListStatus{
			value: "active",
		},
	}
}

func (c ImageListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageListStatus) UnmarshalJSON(b []byte) error {
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

type ImageListDiskFormat struct {
	value string
}

type ImageListDiskFormatEnum struct {
	ZVHD2 ImageListDiskFormat
	VHD   ImageListDiskFormat
	ZVHD  ImageListDiskFormat
	RAW   ImageListDiskFormat
	QCOW2 ImageListDiskFormat
}

func GetImageListDiskFormatEnum() ImageListDiskFormatEnum {
	return ImageListDiskFormatEnum{
		ZVHD2: ImageListDiskFormat{
			value: "zvhd2",
		},
		VHD: ImageListDiskFormat{
			value: "vhd",
		},
		ZVHD: ImageListDiskFormat{
			value: "zvhd",
		},
		RAW: ImageListDiskFormat{
			value: "raw",
		},
		QCOW2: ImageListDiskFormat{
			value: "qcow2",
		},
	}
}

func (c ImageListDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageListDiskFormat) UnmarshalJSON(b []byte) error {
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

type ImageListImagetype struct {
	value string
}

type ImageListImagetypeEnum struct {
	GOLD    ImageListImagetype
	PRIVATE ImageListImagetype
}

func GetImageListImagetypeEnum() ImageListImagetypeEnum {
	return ImageListImagetypeEnum{
		GOLD: ImageListImagetype{
			value: "gold",
		},
		PRIVATE: ImageListImagetype{
			value: "private ",
		},
	}
}

func (c ImageListImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageListImagetype) UnmarshalJSON(b []byte) error {
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

type ImageListOsType struct {
	value string
}

type ImageListOsTypeEnum struct {
	LINUX   ImageListOsType
	WINDOWS ImageListOsType
	OTHER   ImageListOsType
}

func GetImageListOsTypeEnum() ImageListOsTypeEnum {
	return ImageListOsTypeEnum{
		LINUX: ImageListOsType{
			value: "Linux",
		},
		WINDOWS: ImageListOsType{
			value: "Windows",
		},
		OTHER: ImageListOsType{
			value: "Other",
		},
	}
}

func (c ImageListOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageListOsType) UnmarshalJSON(b []byte) error {
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
