package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgencyPolicy 委托中授权的权限策略
type AgencyPolicy struct {

	// 权限类型列表，使用Fabric的功能，可能需要授权一些权限策略，可以在policy_types中加入策略类型来授权权限。FABRIC_COMMON_POLICY：基础通用服务相关权限策略，是委托必需的权限策略；FABRIC_SMN_POLICY：消息通知功能相关权限策略，用来将系统通知消息转发到SMN；FABRIC_LAKEFORMATION_POLICY：LakeFormation功能相关权限策略；FABRIC_AOM_POLICY：AOM功能相关权限策略。
	PolicyType AgencyPolicyPolicyType `json:"policy_type"`

	// 是否是必需的。
	IsRequired bool `json:"is_required"`

	// 是否已开启。
	IsEnable bool `json:"is_enable"`

	// 权限列表
	Actions []string `json:"actions"`

	// 策略描述
	Description string `json:"description"`
}

func (o AgencyPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyPolicy struct{}"
	}

	return strings.Join([]string{"AgencyPolicy", string(data)}, " ")
}

type AgencyPolicyPolicyType struct {
	value string
}

type AgencyPolicyPolicyTypeEnum struct {
	FABRIC_COMMON_POLICY        AgencyPolicyPolicyType
	FABRIC_SMN_POLICY           AgencyPolicyPolicyType
	FABRIC_LAKEFORMATION_POLICY AgencyPolicyPolicyType
	FABRIC_AOM_POLICY           AgencyPolicyPolicyType
}

func GetAgencyPolicyPolicyTypeEnum() AgencyPolicyPolicyTypeEnum {
	return AgencyPolicyPolicyTypeEnum{
		FABRIC_COMMON_POLICY: AgencyPolicyPolicyType{
			value: "FABRIC_COMMON_POLICY",
		},
		FABRIC_SMN_POLICY: AgencyPolicyPolicyType{
			value: "FABRIC_SMN_POLICY",
		},
		FABRIC_LAKEFORMATION_POLICY: AgencyPolicyPolicyType{
			value: "FABRIC_LAKEFORMATION_POLICY",
		},
		FABRIC_AOM_POLICY: AgencyPolicyPolicyType{
			value: "FABRIC_AOM_POLICY",
		},
	}
}

func (c AgencyPolicyPolicyType) Value() string {
	return c.value
}

func (c AgencyPolicyPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyPolicyPolicyType) UnmarshalJSON(b []byte) error {
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
