package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccountAssignmentOperationStatusMetadataDto 提供有关AccountAssignment创建请求的信息
type AccountAssignmentOperationStatusMetadataDto struct {

	// 创建日期
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 请求唯一标识
	RequestId *string `json:"request_id,omitempty"`

	// 权限集授权状态
	Status *AccountAssignmentOperationStatusMetadataDtoStatus `json:"status,omitempty"`
}

func (o AccountAssignmentOperationStatusMetadataDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountAssignmentOperationStatusMetadataDto struct{}"
	}

	return strings.Join([]string{"AccountAssignmentOperationStatusMetadataDto", string(data)}, " ")
}

type AccountAssignmentOperationStatusMetadataDtoStatus struct {
	value string
}

type AccountAssignmentOperationStatusMetadataDtoStatusEnum struct {
	IN_PROGRESS AccountAssignmentOperationStatusMetadataDtoStatus
	FAILED      AccountAssignmentOperationStatusMetadataDtoStatus
	SUCCEEDED   AccountAssignmentOperationStatusMetadataDtoStatus
}

func GetAccountAssignmentOperationStatusMetadataDtoStatusEnum() AccountAssignmentOperationStatusMetadataDtoStatusEnum {
	return AccountAssignmentOperationStatusMetadataDtoStatusEnum{
		IN_PROGRESS: AccountAssignmentOperationStatusMetadataDtoStatus{
			value: "IN_PROGRESS",
		},
		FAILED: AccountAssignmentOperationStatusMetadataDtoStatus{
			value: "FAILED",
		},
		SUCCEEDED: AccountAssignmentOperationStatusMetadataDtoStatus{
			value: "SUCCEEDED",
		},
	}
}

func (c AccountAssignmentOperationStatusMetadataDtoStatus) Value() string {
	return c.value
}

func (c AccountAssignmentOperationStatusMetadataDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountAssignmentOperationStatusMetadataDtoStatus) UnmarshalJSON(b []byte) error {
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
