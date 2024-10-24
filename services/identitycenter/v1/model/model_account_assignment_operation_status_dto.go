package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccountAssignmentOperationStatusDto 授权状态信息
type AccountAssignmentOperationStatusDto struct {

	// 创建日期
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 失败原因
	FailureReason *string `json:"failure_reason,omitempty"`

	// 权限集唯一标识
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// IAM身份中心中的一个实体身份唯一标识，例如用户或用户组
	PrincipalId *string `json:"principal_id,omitempty"`

	// 操作的实体类型
	PrincipalType *AccountAssignmentOperationStatusDtoPrincipalType `json:"principal_type,omitempty"`

	// 请求唯一标识
	RequestId *string `json:"request_id,omitempty"`

	// 权限集授权状态
	Status *AccountAssignmentOperationStatusDtoStatus `json:"status,omitempty"`

	// 目标实体的唯一标识
	TargetId *string `json:"target_id,omitempty"`

	// 实体类型
	TargetType *AccountAssignmentOperationStatusDtoTargetType `json:"target_type,omitempty"`
}

func (o AccountAssignmentOperationStatusDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountAssignmentOperationStatusDto struct{}"
	}

	return strings.Join([]string{"AccountAssignmentOperationStatusDto", string(data)}, " ")
}

type AccountAssignmentOperationStatusDtoPrincipalType struct {
	value string
}

type AccountAssignmentOperationStatusDtoPrincipalTypeEnum struct {
	USER  AccountAssignmentOperationStatusDtoPrincipalType
	GROUP AccountAssignmentOperationStatusDtoPrincipalType
}

func GetAccountAssignmentOperationStatusDtoPrincipalTypeEnum() AccountAssignmentOperationStatusDtoPrincipalTypeEnum {
	return AccountAssignmentOperationStatusDtoPrincipalTypeEnum{
		USER: AccountAssignmentOperationStatusDtoPrincipalType{
			value: "USER",
		},
		GROUP: AccountAssignmentOperationStatusDtoPrincipalType{
			value: "GROUP",
		},
	}
}

func (c AccountAssignmentOperationStatusDtoPrincipalType) Value() string {
	return c.value
}

func (c AccountAssignmentOperationStatusDtoPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountAssignmentOperationStatusDtoPrincipalType) UnmarshalJSON(b []byte) error {
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

type AccountAssignmentOperationStatusDtoStatus struct {
	value string
}

type AccountAssignmentOperationStatusDtoStatusEnum struct {
	IN_PROGRESS AccountAssignmentOperationStatusDtoStatus
	FAILED      AccountAssignmentOperationStatusDtoStatus
	SUCCEEDED   AccountAssignmentOperationStatusDtoStatus
}

func GetAccountAssignmentOperationStatusDtoStatusEnum() AccountAssignmentOperationStatusDtoStatusEnum {
	return AccountAssignmentOperationStatusDtoStatusEnum{
		IN_PROGRESS: AccountAssignmentOperationStatusDtoStatus{
			value: "IN_PROGRESS",
		},
		FAILED: AccountAssignmentOperationStatusDtoStatus{
			value: "FAILED",
		},
		SUCCEEDED: AccountAssignmentOperationStatusDtoStatus{
			value: "SUCCEEDED",
		},
	}
}

func (c AccountAssignmentOperationStatusDtoStatus) Value() string {
	return c.value
}

func (c AccountAssignmentOperationStatusDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountAssignmentOperationStatusDtoStatus) UnmarshalJSON(b []byte) error {
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

type AccountAssignmentOperationStatusDtoTargetType struct {
	value string
}

type AccountAssignmentOperationStatusDtoTargetTypeEnum struct {
	ACCOUNT AccountAssignmentOperationStatusDtoTargetType
}

func GetAccountAssignmentOperationStatusDtoTargetTypeEnum() AccountAssignmentOperationStatusDtoTargetTypeEnum {
	return AccountAssignmentOperationStatusDtoTargetTypeEnum{
		ACCOUNT: AccountAssignmentOperationStatusDtoTargetType{
			value: "ACCOUNT",
		},
	}
}

func (c AccountAssignmentOperationStatusDtoTargetType) Value() string {
	return c.value
}

func (c AccountAssignmentOperationStatusDtoTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountAssignmentOperationStatusDtoTargetType) UnmarshalJSON(b []byte) error {
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
