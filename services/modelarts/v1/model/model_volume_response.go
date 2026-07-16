package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VolumeResponse struct {

	// **参数解释**：notebook返回的扩展存储类型 **参数约束**：不涉及 - OBS：对象存储服务 - OBSFS：并行文件存储 - EFS：弹性文件服务
	Category *VolumeResponseCategory `json:"category,omitempty"`

	// **参数解释**：存储挂载至Notebook实例的目录 **参数约束**：不涉及
	MountPath *string `json:"mount_path,omitempty"`

	// **参数解释**：当category为OBS、OBSFS时，挂载存储源路径。 **参数约束**：不涉及
	Url *string `json:"url,omitempty"`

	// **参数解释**：存储状态 - MOUNTING: 正在挂载中； - MOUNTED: 已挂载成功； - UNMOUNTING: 正在卸载中； - UNMOUNTED: 已卸载完成； - MOUNT_FAILED: 挂载失败 - UNMOUNT_FAILED：卸载失败； **参数约束**：不涉及
	Status *VolumeResponseStatus `json:"status,omitempty"`

	// **参数解释**：存储挂载类型，枚举类。 **约束限制**：无限制。 - STATIC:不支持在实例运行期间挂载以及卸载的存储 - DYNAMIC:支持在实例运行期间挂载以及卸载的存储
	MountType *VolumeResponseMountType `json:"mount_type,omitempty"`
}

func (o VolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeResponse struct{}"
	}

	return strings.Join([]string{"VolumeResponse", string(data)}, " ")
}

type VolumeResponseCategory struct {
	value string
}

type VolumeResponseCategoryEnum struct {
	OBS   VolumeResponseCategory
	OBSFS VolumeResponseCategory
	EFS   VolumeResponseCategory
}

func GetVolumeResponseCategoryEnum() VolumeResponseCategoryEnum {
	return VolumeResponseCategoryEnum{
		OBS: VolumeResponseCategory{
			value: "OBS",
		},
		OBSFS: VolumeResponseCategory{
			value: "OBSFS",
		},
		EFS: VolumeResponseCategory{
			value: "EFS",
		},
	}
}

func (c VolumeResponseCategory) Value() string {
	return c.value
}

func (c VolumeResponseCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeResponseCategory) UnmarshalJSON(b []byte) error {
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

type VolumeResponseStatus struct {
	value string
}

type VolumeResponseStatusEnum struct {
	MOUNTING       VolumeResponseStatus
	MOUNTED        VolumeResponseStatus
	UNMOUNTING     VolumeResponseStatus
	UNMOUNTED      VolumeResponseStatus
	MOUNT_FAILED   VolumeResponseStatus
	UNMOUNT_FAILED VolumeResponseStatus
}

func GetVolumeResponseStatusEnum() VolumeResponseStatusEnum {
	return VolumeResponseStatusEnum{
		MOUNTING: VolumeResponseStatus{
			value: "MOUNTING",
		},
		MOUNTED: VolumeResponseStatus{
			value: "MOUNTED",
		},
		UNMOUNTING: VolumeResponseStatus{
			value: "UNMOUNTING",
		},
		UNMOUNTED: VolumeResponseStatus{
			value: "UNMOUNTED",
		},
		MOUNT_FAILED: VolumeResponseStatus{
			value: "MOUNT_FAILED",
		},
		UNMOUNT_FAILED: VolumeResponseStatus{
			value: "UNMOUNT_FAILED",
		},
	}
}

func (c VolumeResponseStatus) Value() string {
	return c.value
}

func (c VolumeResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeResponseStatus) UnmarshalJSON(b []byte) error {
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

type VolumeResponseMountType struct {
	value string
}

type VolumeResponseMountTypeEnum struct {
	STATIC  VolumeResponseMountType
	DYNAMIC VolumeResponseMountType
}

func GetVolumeResponseMountTypeEnum() VolumeResponseMountTypeEnum {
	return VolumeResponseMountTypeEnum{
		STATIC: VolumeResponseMountType{
			value: "STATIC",
		},
		DYNAMIC: VolumeResponseMountType{
			value: "DYNAMIC",
		},
	}
}

func (c VolumeResponseMountType) Value() string {
	return c.value
}

func (c VolumeResponseMountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeResponseMountType) UnmarshalJSON(b []byte) error {
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
