package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteAgencyRequest Request Object
type DeleteAgencyRequest struct {

	// 权限策略类型。FABRIC_COMMON_POLICY：基础通用服务相关的权限策略；FABRIC_SMN_POLICY：消息通知功能相关的权限策略；FABRIC_LAKEFORMATION_POLICY：LakeFormation功能相关的权限策略；FABRIC_AOM_POLICY：AOM功能相关的权限策略。
	PolicyType *DeleteAgencyRequestPolicyType `json:"policy_type,omitempty"`
}

func (o DeleteAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgencyRequest", string(data)}, " ")
}

type DeleteAgencyRequestPolicyType struct {
	value string
}

type DeleteAgencyRequestPolicyTypeEnum struct {
	FABRIC_COMMON_POLICY        DeleteAgencyRequestPolicyType
	FABRIC_SMN_POLICY           DeleteAgencyRequestPolicyType
	FABRIC_LAKEFORMATION_POLICY DeleteAgencyRequestPolicyType
	FABRIC_AOM_POLICY           DeleteAgencyRequestPolicyType
}

func GetDeleteAgencyRequestPolicyTypeEnum() DeleteAgencyRequestPolicyTypeEnum {
	return DeleteAgencyRequestPolicyTypeEnum{
		FABRIC_COMMON_POLICY: DeleteAgencyRequestPolicyType{
			value: "FABRIC_COMMON_POLICY",
		},
		FABRIC_SMN_POLICY: DeleteAgencyRequestPolicyType{
			value: "FABRIC_SMN_POLICY",
		},
		FABRIC_LAKEFORMATION_POLICY: DeleteAgencyRequestPolicyType{
			value: "FABRIC_LAKEFORMATION_POLICY",
		},
		FABRIC_AOM_POLICY: DeleteAgencyRequestPolicyType{
			value: "FABRIC_AOM_POLICY",
		},
	}
}

func (c DeleteAgencyRequestPolicyType) Value() string {
	return c.value
}

func (c DeleteAgencyRequestPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteAgencyRequestPolicyType) UnmarshalJSON(b []byte) error {
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
