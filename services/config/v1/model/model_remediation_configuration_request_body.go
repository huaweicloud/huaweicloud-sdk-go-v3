package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RemediationConfigurationRequestBody 合规规则修正配置的请求体。
type RemediationConfigurationRequestBody struct {

	// 是否为自动修正。
	Automatic *bool `json:"automatic,omitempty"`

	// 合规规则修正执行的方式。
	TargetType RemediationConfigurationRequestBodyTargetType `json:"target_type"`

	// 修正执行的目标ID。如果修正方式为fgs，则该值为函数工作流的函数urn；如果修正方式为rfs，则该值为资源编排服务的模板name与版本号，两者以/分割，如果没有指定默认V1。
	TargetId string `json:"target_id"`

	// 修正执行的静态参数。
	StaticParameter *[]RemediationStaticParameter `json:"static_parameter,omitempty"`

	ResourceParameter *RemediationResourceParameter `json:"resource_parameter"`

	// 指定时间内修正的最大尝试次数。
	MaximumAttempts *int32 `json:"maximum_attempts,omitempty"`

	// 用于防止循环修正的时间窗口，如果在指定时间内进行了自动修正的最大尝试次数，则将资源添加至修正例外。
	RetryAttemptSeconds *int32 `json:"retry_attempt_seconds,omitempty"`

	// 合规规则修正配置的权限方式。
	AuthType *RemediationConfigurationRequestBodyAuthType `json:"auth_type,omitempty"`

	// 合规规则修正配置的权限信息。
	AuthValue *string `json:"auth_value,omitempty"`
}

func (o RemediationConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"RemediationConfigurationRequestBody", string(data)}, " ")
}

type RemediationConfigurationRequestBodyTargetType struct {
	value string
}

type RemediationConfigurationRequestBodyTargetTypeEnum struct {
	FGS RemediationConfigurationRequestBodyTargetType
	RFS RemediationConfigurationRequestBodyTargetType
}

func GetRemediationConfigurationRequestBodyTargetTypeEnum() RemediationConfigurationRequestBodyTargetTypeEnum {
	return RemediationConfigurationRequestBodyTargetTypeEnum{
		FGS: RemediationConfigurationRequestBodyTargetType{
			value: "fgs",
		},
		RFS: RemediationConfigurationRequestBodyTargetType{
			value: "rfs",
		},
	}
}

func (c RemediationConfigurationRequestBodyTargetType) Value() string {
	return c.value
}

func (c RemediationConfigurationRequestBodyTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RemediationConfigurationRequestBodyTargetType) UnmarshalJSON(b []byte) error {
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

type RemediationConfigurationRequestBodyAuthType struct {
	value string
}

type RemediationConfigurationRequestBodyAuthTypeEnum struct {
	AGENCY       RemediationConfigurationRequestBodyAuthType
	TRUST_AGENCY RemediationConfigurationRequestBodyAuthType
}

func GetRemediationConfigurationRequestBodyAuthTypeEnum() RemediationConfigurationRequestBodyAuthTypeEnum {
	return RemediationConfigurationRequestBodyAuthTypeEnum{
		AGENCY: RemediationConfigurationRequestBodyAuthType{
			value: "agency",
		},
		TRUST_AGENCY: RemediationConfigurationRequestBodyAuthType{
			value: "trustAgency",
		},
	}
}

func (c RemediationConfigurationRequestBodyAuthType) Value() string {
	return c.value
}

func (c RemediationConfigurationRequestBodyAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RemediationConfigurationRequestBodyAuthType) UnmarshalJSON(b []byte) error {
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
