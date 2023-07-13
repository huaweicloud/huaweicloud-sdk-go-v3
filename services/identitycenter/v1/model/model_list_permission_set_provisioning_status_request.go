package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPermissionSetProvisioningStatusRequest Request Object
type ListPermissionSetProvisioningStatusRequest struct {
	InstanceId string `json:"instance_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	// Filters he operation status list based on the passed attribute value.
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
