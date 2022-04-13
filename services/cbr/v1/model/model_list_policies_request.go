package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
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

func (c ListPoliciesRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPoliciesRequestOperationType) UnmarshalJSON(b []byte) error {
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
