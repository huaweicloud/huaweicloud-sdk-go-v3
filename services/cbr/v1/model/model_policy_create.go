package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PolicyCreate struct {
	// 是否启用策略

	Enabled *bool `json:"enabled,omitempty"`
	// 策略名称，长度限制：1- 64，只能由中文、字母、数字、“_”、“-”组成。

	Name string `json:"name"`

	OperationDefinition *PolicyoOdCreate `json:"operation_definition"`
	// 策略类型，如备份，复制 Enum:[ backup，replication]

	OperationType PolicyCreateOperationType `json:"operation_type"`

	Trigger *PolicyTriggerReq `json:"trigger"`
}

func (o PolicyCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyCreate struct{}"
	}

	return strings.Join([]string{"PolicyCreate", string(data)}, " ")
}

type PolicyCreateOperationType struct {
	value string
}

type PolicyCreateOperationTypeEnum struct {
	BACKUP      PolicyCreateOperationType
	REPLICATION PolicyCreateOperationType
}

func GetPolicyCreateOperationTypeEnum() PolicyCreateOperationTypeEnum {
	return PolicyCreateOperationTypeEnum{
		BACKUP: PolicyCreateOperationType{
			value: "backup",
		},
		REPLICATION: PolicyCreateOperationType{
			value: "replication",
		},
	}
}

func (c PolicyCreateOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyCreateOperationType) UnmarshalJSON(b []byte) error {
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
