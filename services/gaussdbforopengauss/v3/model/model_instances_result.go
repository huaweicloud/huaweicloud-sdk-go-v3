package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstancesResult struct {

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 存储类型。
	VolumeType *string `json:"volume_type,omitempty"`

	// 磁盘大小，单位：GB。
	DataVolumeSize float32 `json:"data_volume_size,omitempty"`

	// 实例版本信息。
	Version *string `json:"version,omitempty"`

	// 部署形态。
	Mode *InstancesResultMode `json:"mode,omitempty"`

	// 实例模型，企业版，标准版，基础版。
	InstanceMode *InstancesResultInstanceMode `json:"instance_mode,omitempty"`
}

func (o InstancesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesResult struct{}"
	}

	return strings.Join([]string{"InstancesResult", string(data)}, " ")
}

type InstancesResultMode struct {
	value string
}

type InstancesResultModeEnum struct {
	HA          InstancesResultMode
	INDEPENDENT InstancesResultMode
}

func GetInstancesResultModeEnum() InstancesResultModeEnum {
	return InstancesResultModeEnum{
		HA: InstancesResultMode{
			value: "Ha",
		},
		INDEPENDENT: InstancesResultMode{
			value: "Independent",
		},
	}
}

func (c InstancesResultMode) Value() string {
	return c.value
}

func (c InstancesResultMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstancesResultMode) UnmarshalJSON(b []byte) error {
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

type InstancesResultInstanceMode struct {
	value string
}

type InstancesResultInstanceModeEnum struct {
	ENTERPRISE InstancesResultInstanceMode
	STANDARD   InstancesResultInstanceMode
	BASIC      InstancesResultInstanceMode
}

func GetInstancesResultInstanceModeEnum() InstancesResultInstanceModeEnum {
	return InstancesResultInstanceModeEnum{
		ENTERPRISE: InstancesResultInstanceMode{
			value: "enterprise",
		},
		STANDARD: InstancesResultInstanceMode{
			value: "standard",
		},
		BASIC: InstancesResultInstanceMode{
			value: "basic",
		},
	}
}

func (c InstancesResultInstanceMode) Value() string {
	return c.value
}

func (c InstancesResultInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstancesResultInstanceMode) UnmarshalJSON(b []byte) error {
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
