package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPermissionSetsProvisionedToAccountRequest Request Object
type ListPermissionSetsProvisionedToAccountRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 每个请求返回的最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`

	// 指定账户的唯一身份标识
	AccountId string `json:"account_id"`

	// 权限集授权状态
	ProvisioningStatus *ListPermissionSetsProvisionedToAccountRequestProvisioningStatus `json:"provisioning_status,omitempty"`
}

func (o ListPermissionSetsProvisionedToAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionSetsProvisionedToAccountRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionSetsProvisionedToAccountRequest", string(data)}, " ")
}

type ListPermissionSetsProvisionedToAccountRequestProvisioningStatus struct {
	value string
}

type ListPermissionSetsProvisionedToAccountRequestProvisioningStatusEnum struct {
	LATEST_PERMISSION_SET_PROVISIONED     ListPermissionSetsProvisionedToAccountRequestProvisioningStatus
	LATEST_PERMISSION_SET_NOT_PROVISIONED ListPermissionSetsProvisionedToAccountRequestProvisioningStatus
}

func GetListPermissionSetsProvisionedToAccountRequestProvisioningStatusEnum() ListPermissionSetsProvisionedToAccountRequestProvisioningStatusEnum {
	return ListPermissionSetsProvisionedToAccountRequestProvisioningStatusEnum{
		LATEST_PERMISSION_SET_PROVISIONED: ListPermissionSetsProvisionedToAccountRequestProvisioningStatus{
			value: "LATEST_PERMISSION_SET_PROVISIONED",
		},
		LATEST_PERMISSION_SET_NOT_PROVISIONED: ListPermissionSetsProvisionedToAccountRequestProvisioningStatus{
			value: "LATEST_PERMISSION_SET_NOT_PROVISIONED",
		},
	}
}

func (c ListPermissionSetsProvisionedToAccountRequestProvisioningStatus) Value() string {
	return c.value
}

func (c ListPermissionSetsProvisionedToAccountRequestProvisioningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPermissionSetsProvisionedToAccountRequestProvisioningStatus) UnmarshalJSON(b []byte) error {
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
