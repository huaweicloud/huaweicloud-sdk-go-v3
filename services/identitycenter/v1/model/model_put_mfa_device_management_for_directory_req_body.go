package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PutMfaDeviceManagementForDirectoryReqBody struct {

	// 关联到IAM身份中心实例的身份源的全局唯一标识符（ID）。
	IdentityStoreId string `json:"identity_store_id"`

	// 允许用户自行管理MFA的行为
	UserPermission PutMfaDeviceManagementForDirectoryReqBodyUserPermission `json:"user_permission"`
}

func (o PutMfaDeviceManagementForDirectoryReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutMfaDeviceManagementForDirectoryReqBody struct{}"
	}

	return strings.Join([]string{"PutMfaDeviceManagementForDirectoryReqBody", string(data)}, " ")
}

type PutMfaDeviceManagementForDirectoryReqBodyUserPermission struct {
	value string
}

type PutMfaDeviceManagementForDirectoryReqBodyUserPermissionEnum struct {
	READ_ACTIONS PutMfaDeviceManagementForDirectoryReqBodyUserPermission
	ALL_ACTIONS  PutMfaDeviceManagementForDirectoryReqBodyUserPermission
}

func GetPutMfaDeviceManagementForDirectoryReqBodyUserPermissionEnum() PutMfaDeviceManagementForDirectoryReqBodyUserPermissionEnum {
	return PutMfaDeviceManagementForDirectoryReqBodyUserPermissionEnum{
		READ_ACTIONS: PutMfaDeviceManagementForDirectoryReqBodyUserPermission{
			value: "READ_ACTIONS",
		},
		ALL_ACTIONS: PutMfaDeviceManagementForDirectoryReqBodyUserPermission{
			value: "ALL_ACTIONS",
		},
	}
}

func (c PutMfaDeviceManagementForDirectoryReqBodyUserPermission) Value() string {
	return c.value
}

func (c PutMfaDeviceManagementForDirectoryReqBodyUserPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutMfaDeviceManagementForDirectoryReqBodyUserPermission) UnmarshalJSON(b []byte) error {
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
