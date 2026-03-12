package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDryRunConfigReqBody 更新策略试运行的请求体。
type UpdateDryRunConfigReqBody struct {

	// 试运行策略的类型名称，service_control_policy服务控制策略。
	PolicyType UpdateDryRunConfigReqBodyPolicyType `json:"policy_type"`

	// 根的唯一标识符（ID）。
	RootId string `json:"root_id"`

	// 策略试运行的开启状态。enabled：启用；disabled 禁用。
	Status *UpdateDryRunConfigReqBodyStatus `json:"status,omitempty"`

	// 用户桶名称。
	BucketName *string `json:"bucket_name,omitempty"`

	// 用户桶所在的region。
	RegionId *string `json:"region_id,omitempty"`

	// 用户桶的前缀。
	BucketPrefix *string `json:"bucket_prefix,omitempty"`

	// 委托名称。Organizations服务通过此委托将试运行日志写入用户obs桶
	AgencyName *string `json:"agency_name,omitempty"`
}

func (o UpdateDryRunConfigReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDryRunConfigReqBody struct{}"
	}

	return strings.Join([]string{"UpdateDryRunConfigReqBody", string(data)}, " ")
}

type UpdateDryRunConfigReqBodyPolicyType struct {
	value string
}

type UpdateDryRunConfigReqBodyPolicyTypeEnum struct {
	SERVICE_CONTROL_POLICY UpdateDryRunConfigReqBodyPolicyType
}

func GetUpdateDryRunConfigReqBodyPolicyTypeEnum() UpdateDryRunConfigReqBodyPolicyTypeEnum {
	return UpdateDryRunConfigReqBodyPolicyTypeEnum{
		SERVICE_CONTROL_POLICY: UpdateDryRunConfigReqBodyPolicyType{
			value: "service_control_policy",
		},
	}
}

func (c UpdateDryRunConfigReqBodyPolicyType) Value() string {
	return c.value
}

func (c UpdateDryRunConfigReqBodyPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDryRunConfigReqBodyPolicyType) UnmarshalJSON(b []byte) error {
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

type UpdateDryRunConfigReqBodyStatus struct {
	value string
}

type UpdateDryRunConfigReqBodyStatusEnum struct {
	ENABLED  UpdateDryRunConfigReqBodyStatus
	DISABLED UpdateDryRunConfigReqBodyStatus
}

func GetUpdateDryRunConfigReqBodyStatusEnum() UpdateDryRunConfigReqBodyStatusEnum {
	return UpdateDryRunConfigReqBodyStatusEnum{
		ENABLED: UpdateDryRunConfigReqBodyStatus{
			value: "enabled",
		},
		DISABLED: UpdateDryRunConfigReqBodyStatus{
			value: "disabled",
		},
	}
}

func (c UpdateDryRunConfigReqBodyStatus) Value() string {
	return c.value
}

func (c UpdateDryRunConfigReqBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDryRunConfigReqBodyStatus) UnmarshalJSON(b []byte) error {
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
