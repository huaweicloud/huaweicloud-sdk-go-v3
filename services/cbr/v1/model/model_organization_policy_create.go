package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OrganizationPolicyCreate 创建组织策略body体
type OrganizationPolicyCreate struct {

	// 组织策略名称
	Name string `json:"name"`

	// 组织策略描述
	Description *string `json:"description,omitempty"`

	// 组织策略类型 - backup: 备份 - replication: 复制
	OperationType OrganizationPolicyCreateOperationType `json:"operation_type"`

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 策略是否开启
	PolicyEnabled bool `json:"policy_enabled"`

	PolicyOperationDefinition *PolicyoOdCreate `json:"policy_operation_definition"`

	PolicyTrigger *PolicyTriggerReq `json:"policy_trigger"`
}

func (o OrganizationPolicyCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicyCreate struct{}"
	}

	return strings.Join([]string{"OrganizationPolicyCreate", string(data)}, " ")
}

type OrganizationPolicyCreateOperationType struct {
	value string
}

type OrganizationPolicyCreateOperationTypeEnum struct {
	BACKUP      OrganizationPolicyCreateOperationType
	REPLICATION OrganizationPolicyCreateOperationType
}

func GetOrganizationPolicyCreateOperationTypeEnum() OrganizationPolicyCreateOperationTypeEnum {
	return OrganizationPolicyCreateOperationTypeEnum{
		BACKUP: OrganizationPolicyCreateOperationType{
			value: "backup",
		},
		REPLICATION: OrganizationPolicyCreateOperationType{
			value: "replication",
		},
	}
}

func (c OrganizationPolicyCreateOperationType) Value() string {
	return c.value
}

func (c OrganizationPolicyCreateOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrganizationPolicyCreateOperationType) UnmarshalJSON(b []byte) error {
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
