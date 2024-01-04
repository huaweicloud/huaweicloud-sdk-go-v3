package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMonitorSystemResponseBodySpec 查询监控系统详情。
type ShowMonitorSystemResponseBodySpec struct {

	// 监控系统id。
	Id *string `json:"id,omitempty"`

	// 采集方式。
	Type *ShowMonitorSystemResponseBodySpecType `json:"type,omitempty"`

	// 探针注入方式。
	Instrumentation *ShowMonitorSystemResponseBodySpecInstrumentation `json:"instrumentation,omitempty"`

	// apm2访问密钥Key。
	AccessKey *string `json:"access_key,omitempty"`

	// apm2访问密钥value。
	AccessValue *string `json:"access_value,omitempty"`

	// apm2 opentelemetry接入token。
	AccessToken *string `json:"access_token,omitempty"`

	// 探针接入点。
	MasterAddress *string `json:"master_address,omitempty"`

	// apm2应用。
	ApmApplication *string `json:"apm_application,omitempty"`

	// 增强型探针/opentelemetry探针版本。
	Version *string `json:"version,omitempty"`

	// 探针镜像更新策略。
	ImagePullPolicy *ShowMonitorSystemResponseBodySpecImagePullPolicy `json:"image_pull_policy,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ShowMonitorSystemResponseBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitorSystemResponseBodySpec struct{}"
	}

	return strings.Join([]string{"ShowMonitorSystemResponseBodySpec", string(data)}, " ")
}

type ShowMonitorSystemResponseBodySpecType struct {
	value string
}

type ShowMonitorSystemResponseBodySpecTypeEnum struct {
	OPENTELEMETRY ShowMonitorSystemResponseBodySpecType
	APM2          ShowMonitorSystemResponseBodySpecType
}

func GetShowMonitorSystemResponseBodySpecTypeEnum() ShowMonitorSystemResponseBodySpecTypeEnum {
	return ShowMonitorSystemResponseBodySpecTypeEnum{
		OPENTELEMETRY: ShowMonitorSystemResponseBodySpecType{
			value: "opentelemetry",
		},
		APM2: ShowMonitorSystemResponseBodySpecType{
			value: "apm2",
		},
	}
}

func (c ShowMonitorSystemResponseBodySpecType) Value() string {
	return c.value
}

func (c ShowMonitorSystemResponseBodySpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMonitorSystemResponseBodySpecType) UnmarshalJSON(b []byte) error {
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

type ShowMonitorSystemResponseBodySpecInstrumentation struct {
	value string
}

type ShowMonitorSystemResponseBodySpecInstrumentationEnum struct {
	MANUAL    ShowMonitorSystemResponseBodySpecInstrumentation
	AUTOMATIC ShowMonitorSystemResponseBodySpecInstrumentation
}

func GetShowMonitorSystemResponseBodySpecInstrumentationEnum() ShowMonitorSystemResponseBodySpecInstrumentationEnum {
	return ShowMonitorSystemResponseBodySpecInstrumentationEnum{
		MANUAL: ShowMonitorSystemResponseBodySpecInstrumentation{
			value: "manual",
		},
		AUTOMATIC: ShowMonitorSystemResponseBodySpecInstrumentation{
			value: "automatic",
		},
	}
}

func (c ShowMonitorSystemResponseBodySpecInstrumentation) Value() string {
	return c.value
}

func (c ShowMonitorSystemResponseBodySpecInstrumentation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMonitorSystemResponseBodySpecInstrumentation) UnmarshalJSON(b []byte) error {
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

type ShowMonitorSystemResponseBodySpecImagePullPolicy struct {
	value string
}

type ShowMonitorSystemResponseBodySpecImagePullPolicyEnum struct {
	ALWAYS         ShowMonitorSystemResponseBodySpecImagePullPolicy
	IF_NOT_PRESENT ShowMonitorSystemResponseBodySpecImagePullPolicy
}

func GetShowMonitorSystemResponseBodySpecImagePullPolicyEnum() ShowMonitorSystemResponseBodySpecImagePullPolicyEnum {
	return ShowMonitorSystemResponseBodySpecImagePullPolicyEnum{
		ALWAYS: ShowMonitorSystemResponseBodySpecImagePullPolicy{
			value: "Always",
		},
		IF_NOT_PRESENT: ShowMonitorSystemResponseBodySpecImagePullPolicy{
			value: "IfNotPresent",
		},
	}
}

func (c ShowMonitorSystemResponseBodySpecImagePullPolicy) Value() string {
	return c.value
}

func (c ShowMonitorSystemResponseBodySpecImagePullPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMonitorSystemResponseBodySpecImagePullPolicy) UnmarshalJSON(b []byte) error {
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
