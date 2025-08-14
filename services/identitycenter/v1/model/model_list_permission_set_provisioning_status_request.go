package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPermissionSetProvisioningStatusRequest Request Object
type ListPermissionSetProvisioningStatusRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 每个请求返回的最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`

	// 根据传递的属性值过滤操作状态列表
	Status *ListPermissionSetProvisioningStatusRequestStatus `json:"status,omitempty"`
}

func (o ListPermissionSetProvisioningStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionSetProvisioningStatusRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionSetProvisioningStatusRequest", string(data)}, " ")
}

type ListPermissionSetProvisioningStatusRequestStatus struct {
	value string
}

type ListPermissionSetProvisioningStatusRequestStatusEnum struct {
	IN_PROGRESS ListPermissionSetProvisioningStatusRequestStatus
	SUCCEEDED   ListPermissionSetProvisioningStatusRequestStatus
	FAILED      ListPermissionSetProvisioningStatusRequestStatus
}

func GetListPermissionSetProvisioningStatusRequestStatusEnum() ListPermissionSetProvisioningStatusRequestStatusEnum {
	return ListPermissionSetProvisioningStatusRequestStatusEnum{
		IN_PROGRESS: ListPermissionSetProvisioningStatusRequestStatus{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: ListPermissionSetProvisioningStatusRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListPermissionSetProvisioningStatusRequestStatus{
			value: "FAILED",
		},
	}
}

func (c ListPermissionSetProvisioningStatusRequestStatus) Value() string {
	return c.value
}

func (c ListPermissionSetProvisioningStatusRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPermissionSetProvisioningStatusRequestStatus) UnmarshalJSON(b []byte) error {
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
