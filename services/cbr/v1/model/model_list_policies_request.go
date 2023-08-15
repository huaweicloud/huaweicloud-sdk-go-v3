package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPoliciesRequest Request Object
type ListPoliciesRequest struct {

	// 策略类型：备份（backup）、复制(replication)
	OperationType *ListPoliciesRequestOperationType `json:"operation_type,omitempty"`

	// 存储库ID
	VaultId *string `json:"vault_id,omitempty"`
}

func (o ListPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListPoliciesRequest", string(data)}, " ")
}

type ListPoliciesRequestOperationType struct {
	value string
}

type ListPoliciesRequestOperationTypeEnum struct {
	BACKUP      ListPoliciesRequestOperationType
	REPLICATION ListPoliciesRequestOperationType
}

func GetListPoliciesRequestOperationTypeEnum() ListPoliciesRequestOperationTypeEnum {
	return ListPoliciesRequestOperationTypeEnum{
		BACKUP: ListPoliciesRequestOperationType{
			value: "backup",
		},
		REPLICATION: ListPoliciesRequestOperationType{
			value: "replication",
		},
	}
}

func (c ListPoliciesRequestOperationType) Value() string {
	return c.value
}

func (c ListPoliciesRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPoliciesRequestOperationType) UnmarshalJSON(b []byte) error {
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
