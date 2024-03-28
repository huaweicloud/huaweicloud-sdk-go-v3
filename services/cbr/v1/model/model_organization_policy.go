package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OrganizationPolicy 组织策略
type OrganizationPolicy struct {

	// 组织策略ID
	Id string `json:"id"`

	// 组织策略名称
	Name string `json:"name"`

	// 组织策略描述
	Description *string `json:"description,omitempty"`

	// 组织策略类型 - backup: 备份 -  replication: 复制
	OperationType OrganizationPolicyOperationType `json:"operation_type"`

	// 组织策略所属账号ID
	DomainId string `json:"domain_id"`

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 策略是否开启
	PolicyEnabled bool `json:"policy_enabled"`

	PolicyOperationDefinition *PolicyoOdCreate `json:"policy_operation_definition"`

	PolicyTrigger *PolicyTriggerReq `json:"policy_trigger"`

	// 组织策略状态
	Status string `json:"status"`

	// 组织策略所属账号
	DomainName *string `json:"domain_name,omitempty"`
}

func (o OrganizationPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicy struct{}"
	}

	return strings.Join([]string{"OrganizationPolicy", string(data)}, " ")
}

type OrganizationPolicyOperationType struct {
	value string
}

type OrganizationPolicyOperationTypeEnum struct {
	BACKUP      OrganizationPolicyOperationType
	REPLICATION OrganizationPolicyOperationType
}

func GetOrganizationPolicyOperationTypeEnum() OrganizationPolicyOperationTypeEnum {
	return OrganizationPolicyOperationTypeEnum{
		BACKUP: OrganizationPolicyOperationType{
			value: "backup",
		},
		REPLICATION: OrganizationPolicyOperationType{
			value: " replication",
		},
	}
}

func (c OrganizationPolicyOperationType) Value() string {
	return c.value
}

func (c OrganizationPolicyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrganizationPolicyOperationType) UnmarshalJSON(b []byte) error {
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
