package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Policy struct {

	// 策略是否启用
	Enabled bool `json:"enabled"`

	// 策略ID
	Id string `json:"id"`

	// 策略名称
	Name string `json:"name"`

	OperationDefinition *PolicyoOdCreate `json:"operation_definition"`

	// 策略类型,例如 ‘backup’：自动备份
	OperationType PolicyOperationType `json:"operation_type"`

	Trigger *PolicyTriggerResp `json:"trigger"`

	// 关联的存储库
	AssociatedVaults *[]PolicyAssociateVault `json:"associated_vaults,omitempty"`

	// 策略类型，取值范围如下： - custom_policy：普通用户策略 - organization_policy_v1：老版本组织策略 - organization_policy_v2：新版本组织策略 - organization_policy_removed：退出组织的用户已绑定在存储库上的新版本组织策略（不自动解绑策略与存储库绑定关系，但是不能新绑定存储库）
	PolicyType *string `json:"policy_type,omitempty"`
}

func (o Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policy struct{}"
	}

	return strings.Join([]string{"Policy", string(data)}, " ")
}

type PolicyOperationType struct {
	value string
}

type PolicyOperationTypeEnum struct {
	BACKUP      PolicyOperationType
	REPLICATION PolicyOperationType
}

func GetPolicyOperationTypeEnum() PolicyOperationTypeEnum {
	return PolicyOperationTypeEnum{
		BACKUP: PolicyOperationType{
			value: "backup",
		},
		REPLICATION: PolicyOperationType{
			value: "replication",
		},
	}
}

func (c PolicyOperationType) Value() string {
	return c.value
}

func (c PolicyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyOperationType) UnmarshalJSON(b []byte) error {
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
