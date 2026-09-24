package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateInternalEndpointPermissionsResponse Response Object
type UpdateInternalEndpointPermissionsResponse struct {

	// 权限列表
	Permissions *[]string `json:"permissions,omitempty"`

	// 权限类型 取值范围： - domainId：基于账户ID配置终端节点服务白名单 - orgPath：基于账户所在组织路径配置终端节点服务白名单
	PermissionType *UpdateInternalEndpointPermissionsResponsePermissionType `json:"permission_type,omitempty"`
	HttpStatusCode int                                                      `json:"-"`
}

func (o UpdateInternalEndpointPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternalEndpointPermissionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateInternalEndpointPermissionsResponse", string(data)}, " ")
}

type UpdateInternalEndpointPermissionsResponsePermissionType struct {
	value string
}

type UpdateInternalEndpointPermissionsResponsePermissionTypeEnum struct {
	DOMAIN_ID UpdateInternalEndpointPermissionsResponsePermissionType
	ORG_PATH  UpdateInternalEndpointPermissionsResponsePermissionType
}

func GetUpdateInternalEndpointPermissionsResponsePermissionTypeEnum() UpdateInternalEndpointPermissionsResponsePermissionTypeEnum {
	return UpdateInternalEndpointPermissionsResponsePermissionTypeEnum{
		DOMAIN_ID: UpdateInternalEndpointPermissionsResponsePermissionType{
			value: "domainId",
		},
		ORG_PATH: UpdateInternalEndpointPermissionsResponsePermissionType{
			value: "orgPath",
		},
	}
}

func (c UpdateInternalEndpointPermissionsResponsePermissionType) Value() string {
	return c.value
}

func (c UpdateInternalEndpointPermissionsResponsePermissionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInternalEndpointPermissionsResponsePermissionType) UnmarshalJSON(b []byte) error {
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
