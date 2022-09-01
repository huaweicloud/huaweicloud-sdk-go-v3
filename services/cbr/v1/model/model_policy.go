package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Policy struct {

	// 策略是否启用
	Enabled bool `json:"enabled" xml:"enabled"`

	// 策略ID
	Id string `json:"id" xml:"id"`

	// 策略名称
	Name string `json:"name" xml:"name"`

	OperationDefinition *PolicyoOdCreate `json:"operation_definition" xml:"operation_definition"`

	// 策略类型,例如 ‘backup’：自动备份
	OperationType PolicyOperationType `json:"operation_type" xml:"operation_type"`

	Trigger *PolicyTriggerResp `json:"trigger" xml:"trigger"`

	// 关联的存储库
	AssociatedVaults *[]PolicyAssociateVault `json:"associated_vaults,omitempty" xml:"associated_vaults"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
