package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PermissionSetProvisioningStatusMetadataDto 权限集授权状态信息
type PermissionSetProvisioningStatusMetadataDto struct {

	// 权限集创建日期
	CreatedDate *int64 `json:"created_date,omitempty"`

	// 请求唯一标识
	RequestId *string `json:"request_id,omitempty"`

	// 权限集授权状态
	Status *PermissionSetProvisioningStatusMetadataDtoStatus `json:"status,omitempty"`
}

func (o PermissionSetProvisioningStatusMetadataDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetProvisioningStatusMetadataDto struct{}"
	}

	return strings.Join([]string{"PermissionSetProvisioningStatusMetadataDto", string(data)}, " ")
}

type PermissionSetProvisioningStatusMetadataDtoStatus struct {
	value string
}

type PermissionSetProvisioningStatusMetadataDtoStatusEnum struct {
	IN_PROGRESS PermissionSetProvisioningStatusMetadataDtoStatus
	FAILED      PermissionSetProvisioningStatusMetadataDtoStatus
	SUCCEEDED   PermissionSetProvisioningStatusMetadataDtoStatus
}

func GetPermissionSetProvisioningStatusMetadataDtoStatusEnum() PermissionSetProvisioningStatusMetadataDtoStatusEnum {
	return PermissionSetProvisioningStatusMetadataDtoStatusEnum{
		IN_PROGRESS: PermissionSetProvisioningStatusMetadataDtoStatus{
			value: "IN_PROGRESS",
		},
		FAILED: PermissionSetProvisioningStatusMetadataDtoStatus{
			value: "FAILED",
		},
		SUCCEEDED: PermissionSetProvisioningStatusMetadataDtoStatus{
			value: "SUCCEEDED",
		},
	}
}

func (c PermissionSetProvisioningStatusMetadataDtoStatus) Value() string {
	return c.value
}

func (c PermissionSetProvisioningStatusMetadataDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetProvisioningStatusMetadataDtoStatus) UnmarshalJSON(b []byte) error {
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
