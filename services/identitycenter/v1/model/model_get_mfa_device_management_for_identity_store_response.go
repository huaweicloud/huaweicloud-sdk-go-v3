package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetMfaDeviceManagementForIdentityStoreResponse Response Object
type GetMfaDeviceManagementForIdentityStoreResponse struct {

	// 关联到IAM身份中心实例的身份源的全局唯一标识符（ID）。
	IdentityStoreId *string `json:"identity_store_id,omitempty"`

	// 允许用户自行操作的MFA行为
	UserPermission *GetMfaDeviceManagementForIdentityStoreResponseUserPermission `json:"user_permission,omitempty"`
	HttpStatusCode int                                                           `json:"-"`
}

func (o GetMfaDeviceManagementForIdentityStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetMfaDeviceManagementForIdentityStoreResponse struct{}"
	}

	return strings.Join([]string{"GetMfaDeviceManagementForIdentityStoreResponse", string(data)}, " ")
}

type GetMfaDeviceManagementForIdentityStoreResponseUserPermission struct {
	value string
}

type GetMfaDeviceManagementForIdentityStoreResponseUserPermissionEnum struct {
	READ_ACTIONS GetMfaDeviceManagementForIdentityStoreResponseUserPermission
	ALL_ACTIONS  GetMfaDeviceManagementForIdentityStoreResponseUserPermission
}

func GetGetMfaDeviceManagementForIdentityStoreResponseUserPermissionEnum() GetMfaDeviceManagementForIdentityStoreResponseUserPermissionEnum {
	return GetMfaDeviceManagementForIdentityStoreResponseUserPermissionEnum{
		READ_ACTIONS: GetMfaDeviceManagementForIdentityStoreResponseUserPermission{
			value: "READ_ACTIONS",
		},
		ALL_ACTIONS: GetMfaDeviceManagementForIdentityStoreResponseUserPermission{
			value: "ALL_ACTIONS",
		},
	}
}

func (c GetMfaDeviceManagementForIdentityStoreResponseUserPermission) Value() string {
	return c.value
}

func (c GetMfaDeviceManagementForIdentityStoreResponseUserPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetMfaDeviceManagementForIdentityStoreResponseUserPermission) UnmarshalJSON(b []byte) error {
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
