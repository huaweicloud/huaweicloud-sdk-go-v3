package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PermissionSetProvisioningStatusDto 为指定的权限集提供配置操作的状态.
type PermissionSetProvisioningStatusDto struct {

	// 指定账户的唯一身份标识.
	AccountId *string `json:"account_id,omitempty"`

	// 权限集创建日期
	CreatedDate *string `json:"created_date,omitempty"`

	// 失败原因
	FailureReason *string `json:"failure_reason,omitempty"`

	// 权限集唯一标识
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 请求唯一标识
	RequestId *string `json:"request_id,omitempty"`

	// 权限集授权状态
	Status *PermissionSetProvisioningStatusDtoStatus `json:"status,omitempty"`
}

func (o PermissionSetProvisioningStatusDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetProvisioningStatusDto struct{}"
	}

	return strings.Join([]string{"PermissionSetProvisioningStatusDto", string(data)}, " ")
}

type PermissionSetProvisioningStatusDtoStatus struct {
	value string
}

type PermissionSetProvisioningStatusDtoStatusEnum struct {
	IN_PROGRESS PermissionSetProvisioningStatusDtoStatus
	FAILED      PermissionSetProvisioningStatusDtoStatus
	SUCCEEDED   PermissionSetProvisioningStatusDtoStatus
}

func GetPermissionSetProvisioningStatusDtoStatusEnum() PermissionSetProvisioningStatusDtoStatusEnum {
	return PermissionSetProvisioningStatusDtoStatusEnum{
		IN_PROGRESS: PermissionSetProvisioningStatusDtoStatus{
			value: "IN_PROGRESS",
		},
		FAILED: PermissionSetProvisioningStatusDtoStatus{
			value: "FAILED",
		},
		SUCCEEDED: PermissionSetProvisioningStatusDtoStatus{
			value: "SUCCEEDED",
		},
	}
}

func (c PermissionSetProvisioningStatusDtoStatus) Value() string {
	return c.value
}

func (c PermissionSetProvisioningStatusDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetProvisioningStatusDtoStatus) UnmarshalJSON(b []byte) error {
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
