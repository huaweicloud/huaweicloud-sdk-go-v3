package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListOrganizationPolicyRequest Request Object
type ListOrganizationPolicyRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 分页limit
	Limit *int32 `json:"limit,omitempty"`

	// 分页offset
	Offset *int64 `json:"offset,omitempty"`

	// 组织策略类型，取值范围：backup：备份，replication：复制
	OperationType *ListOrganizationPolicyRequestOperationType `json:"operation_type,omitempty"`
}

func (o ListOrganizationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListOrganizationPolicyRequest", string(data)}, " ")
}

type ListOrganizationPolicyRequestOperationType struct {
	value string
}

type ListOrganizationPolicyRequestOperationTypeEnum struct {
	BACKUP      ListOrganizationPolicyRequestOperationType
	REPLICATION ListOrganizationPolicyRequestOperationType
}

func GetListOrganizationPolicyRequestOperationTypeEnum() ListOrganizationPolicyRequestOperationTypeEnum {
	return ListOrganizationPolicyRequestOperationTypeEnum{
		BACKUP: ListOrganizationPolicyRequestOperationType{
			value: "backup",
		},
		REPLICATION: ListOrganizationPolicyRequestOperationType{
			value: "replication",
		},
	}
}

func (c ListOrganizationPolicyRequestOperationType) Value() string {
	return c.value
}

func (c ListOrganizationPolicyRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOrganizationPolicyRequestOperationType) UnmarshalJSON(b []byte) error {
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
