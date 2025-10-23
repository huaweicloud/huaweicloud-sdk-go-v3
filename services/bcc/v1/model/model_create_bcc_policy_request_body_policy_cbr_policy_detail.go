package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateBccPolicyRequestBodyPolicyCbrPolicyDetail cbr策略详情
type CreateBccPolicyRequestBodyPolicyCbrPolicyDetail struct {
	OperationDefinition *CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationDefinition `json:"operation_definition"`

	// 保护类型,本地备份backup，异地复制replication
	OperationType CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType `json:"operation_type"`

	Trigger *CreateBccPolicyRequestBodyPolicyCbrPolicyDetailTrigger `json:"trigger"`
}

func (o CreateBccPolicyRequestBodyPolicyCbrPolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBccPolicyRequestBodyPolicyCbrPolicyDetail struct{}"
	}

	return strings.Join([]string{"CreateBccPolicyRequestBodyPolicyCbrPolicyDetail", string(data)}, " ")
}

type CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType struct {
	value string
}

type CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationTypeEnum struct {
	BACKUP      CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType
	REPLICATION CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType
}

func GetCreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationTypeEnum() CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationTypeEnum {
	return CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationTypeEnum{
		BACKUP: CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType{
			value: "backup",
		},
		REPLICATION: CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType{
			value: "replication",
		},
	}
}

func (c CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType) Value() string {
	return c.value
}

func (c CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBccPolicyRequestBodyPolicyCbrPolicyDetailOperationType) UnmarshalJSON(b []byte) error {
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
