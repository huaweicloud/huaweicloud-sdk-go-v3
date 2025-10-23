package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPolicyRequest Request Object
type ListPolicyRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 策略类型，备份 backup, 复制 replication
	OperationType *ListPolicyRequestOperationType `json:"operation_type,omitempty"`

	// 分页参数limit
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数offset
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyRequest", string(data)}, " ")
}

type ListPolicyRequestOperationType struct {
	value string
}

type ListPolicyRequestOperationTypeEnum struct {
	BACKUP      ListPolicyRequestOperationType
	REPLICATION ListPolicyRequestOperationType
}

func GetListPolicyRequestOperationTypeEnum() ListPolicyRequestOperationTypeEnum {
	return ListPolicyRequestOperationTypeEnum{
		BACKUP: ListPolicyRequestOperationType{
			value: "backup",
		},
		REPLICATION: ListPolicyRequestOperationType{
			value: "replication",
		},
	}
}

func (c ListPolicyRequestOperationType) Value() string {
	return c.value
}

func (c ListPolicyRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPolicyRequestOperationType) UnmarshalJSON(b []byte) error {
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
