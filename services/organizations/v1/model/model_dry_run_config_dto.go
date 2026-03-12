package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// DryRunConfigDto 策略试运行的详细信息。
type DryRunConfigDto struct {

	// 根的唯一标识符（ID）。
	RootId string `json:"root_id"`

	// 策略试运行的开启状态。enabled：启用；pending_enable：启用中；disabled：禁用；pending_disable：禁用中。
	Status DryRunConfigDtoStatus `json:"status"`

	// 用户桶名称。
	BucketName *string `json:"bucket_name,omitempty"`

	// 用户桶所在的region。
	RegionId *string `json:"region_id,omitempty"`

	// 用户桶的前缀。
	BucketPrefix *string `json:"bucket_prefix,omitempty"`

	// 委托名称。Organizations服务通过此委托将试运行日志写入用户obs桶
	AgencyName *string `json:"agency_name,omitempty"`

	// 试运行策略的类型名称，service_control_policy服务控制策略。
	PolicyType DryRunConfigDtoPolicyType `json:"policy_type"`

	// 策略试运行配置的开启时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 策略试运行配置的更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o DryRunConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DryRunConfigDto struct{}"
	}

	return strings.Join([]string{"DryRunConfigDto", string(data)}, " ")
}

type DryRunConfigDtoStatus struct {
	value string
}

type DryRunConfigDtoStatusEnum struct {
	ENABLED         DryRunConfigDtoStatus
	PENDING_ENABLE  DryRunConfigDtoStatus
	DISABLED        DryRunConfigDtoStatus
	PENDING_DISABLE DryRunConfigDtoStatus
}

func GetDryRunConfigDtoStatusEnum() DryRunConfigDtoStatusEnum {
	return DryRunConfigDtoStatusEnum{
		ENABLED: DryRunConfigDtoStatus{
			value: "enabled",
		},
		PENDING_ENABLE: DryRunConfigDtoStatus{
			value: "pending_enable",
		},
		DISABLED: DryRunConfigDtoStatus{
			value: "disabled",
		},
		PENDING_DISABLE: DryRunConfigDtoStatus{
			value: "pending_disable",
		},
	}
}

func (c DryRunConfigDtoStatus) Value() string {
	return c.value
}

func (c DryRunConfigDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DryRunConfigDtoStatus) UnmarshalJSON(b []byte) error {
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

type DryRunConfigDtoPolicyType struct {
	value string
}

type DryRunConfigDtoPolicyTypeEnum struct {
	SERVICE_CONTROL_POLICY DryRunConfigDtoPolicyType
}

func GetDryRunConfigDtoPolicyTypeEnum() DryRunConfigDtoPolicyTypeEnum {
	return DryRunConfigDtoPolicyTypeEnum{
		SERVICE_CONTROL_POLICY: DryRunConfigDtoPolicyType{
			value: "service_control_policy",
		},
	}
}

func (c DryRunConfigDtoPolicyType) Value() string {
	return c.value
}

func (c DryRunConfigDtoPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DryRunConfigDtoPolicyType) UnmarshalJSON(b []byte) error {
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
