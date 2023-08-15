package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAccountsForProvisionedPermissionSetRequest Request Object
type ListAccountsForProvisionedPermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	// 权限集分配状态.
	ProvisioningStatus *ListAccountsForProvisionedPermissionSetRequestProvisioningStatus `json:"provisioning_status,omitempty"`
}

func (o ListAccountsForProvisionedPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountsForProvisionedPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"ListAccountsForProvisionedPermissionSetRequest", string(data)}, " ")
}

type ListAccountsForProvisionedPermissionSetRequestProvisioningStatus struct {
	value string
}

type ListAccountsForProvisionedPermissionSetRequestProvisioningStatusEnum struct {
	LATEST_PERMISSION_SET_PROVISIONED     ListAccountsForProvisionedPermissionSetRequestProvisioningStatus
	LATEST_PERMISSION_SET_NOT_PROVISIONED ListAccountsForProvisionedPermissionSetRequestProvisioningStatus
}

func GetListAccountsForProvisionedPermissionSetRequestProvisioningStatusEnum() ListAccountsForProvisionedPermissionSetRequestProvisioningStatusEnum {
	return ListAccountsForProvisionedPermissionSetRequestProvisioningStatusEnum{
		LATEST_PERMISSION_SET_PROVISIONED: ListAccountsForProvisionedPermissionSetRequestProvisioningStatus{
			value: "LATEST_PERMISSION_SET_PROVISIONED",
		},
		LATEST_PERMISSION_SET_NOT_PROVISIONED: ListAccountsForProvisionedPermissionSetRequestProvisioningStatus{
			value: "LATEST_PERMISSION_SET_NOT_PROVISIONED",
		},
	}
}

func (c ListAccountsForProvisionedPermissionSetRequestProvisioningStatus) Value() string {
	return c.value
}

func (c ListAccountsForProvisionedPermissionSetRequestProvisioningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccountsForProvisionedPermissionSetRequestProvisioningStatus) UnmarshalJSON(b []byte) error {
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
