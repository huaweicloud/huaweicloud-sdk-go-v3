package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgencyTypeBody 创建委托请求参数
type AgencyTypeBody struct {

	// 权限类型列表，使用Fabric的功能，可能需要授权一些权限策略，可以在policy_types中加入策略类型来授权权限。FABRIC_COMMON_POLICY：基础通用服务权限策略，是委托必需的权限策略；FABRIC_SMN_POLICY：消息通知功能相关权限策略，用来将系统通知消息转发到SMN；FABRIC_LAKEFORMATION_POLICY：LakeFormation功能相关权限策略；FABRIC_AOM_POLICY：AOM功能相关权限策略。
	PolicyTypes *[]AgencyTypeBodyPolicyTypes `json:"policy_types,omitempty"`
}

func (o AgencyTypeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyTypeBody struct{}"
	}

	return strings.Join([]string{"AgencyTypeBody", string(data)}, " ")
}

type AgencyTypeBodyPolicyTypes struct {
	value string
}

type AgencyTypeBodyPolicyTypesEnum struct {
	FABRIC_COMMON_POLICY        AgencyTypeBodyPolicyTypes
	FABRIC_SMN_POLICY           AgencyTypeBodyPolicyTypes
	FABRIC_LAKEFORMATION_POLICY AgencyTypeBodyPolicyTypes
	FABRIC_AOM_POLICY           AgencyTypeBodyPolicyTypes
}

func GetAgencyTypeBodyPolicyTypesEnum() AgencyTypeBodyPolicyTypesEnum {
	return AgencyTypeBodyPolicyTypesEnum{
		FABRIC_COMMON_POLICY: AgencyTypeBodyPolicyTypes{
			value: "FABRIC_COMMON_POLICY",
		},
		FABRIC_SMN_POLICY: AgencyTypeBodyPolicyTypes{
			value: "FABRIC_SMN_POLICY",
		},
		FABRIC_LAKEFORMATION_POLICY: AgencyTypeBodyPolicyTypes{
			value: "FABRIC_LAKEFORMATION_POLICY",
		},
		FABRIC_AOM_POLICY: AgencyTypeBodyPolicyTypes{
			value: "FABRIC_AOM_POLICY",
		},
	}
}

func (c AgencyTypeBodyPolicyTypes) Value() string {
	return c.value
}

func (c AgencyTypeBodyPolicyTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyTypeBodyPolicyTypes) UnmarshalJSON(b []byte) error {
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
