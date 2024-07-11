package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EnvironmentPermissionV2Body 环境修改权限请求体
type EnvironmentPermissionV2Body struct {

	// 角色id
	RoleId *string `json:"role_id,omitempty"`

	// 权限名称，can_view：查看权限；can_edit：编辑权限；can_delete：删除权限；can_deploy：部署权限；can_manage：权限管理权限
	PermissionName *EnvironmentPermissionV2BodyPermissionName `json:"permission_name,omitempty"`

	// true 有权限，false 无权限
	PermissionValue *bool `json:"permission_value,omitempty"`
}

func (o EnvironmentPermissionV2Body) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentPermissionV2Body struct{}"
	}

	return strings.Join([]string{"EnvironmentPermissionV2Body", string(data)}, " ")
}

type EnvironmentPermissionV2BodyPermissionName struct {
	value string
}

type EnvironmentPermissionV2BodyPermissionNameEnum struct {
	CAN_VIEW   EnvironmentPermissionV2BodyPermissionName
	CAN_EDIT   EnvironmentPermissionV2BodyPermissionName
	CAN_DELETE EnvironmentPermissionV2BodyPermissionName
	CAN_DEPLOY EnvironmentPermissionV2BodyPermissionName
	CAN_MANAGE EnvironmentPermissionV2BodyPermissionName
}

func GetEnvironmentPermissionV2BodyPermissionNameEnum() EnvironmentPermissionV2BodyPermissionNameEnum {
	return EnvironmentPermissionV2BodyPermissionNameEnum{
		CAN_VIEW: EnvironmentPermissionV2BodyPermissionName{
			value: "can_view",
		},
		CAN_EDIT: EnvironmentPermissionV2BodyPermissionName{
			value: "can_edit",
		},
		CAN_DELETE: EnvironmentPermissionV2BodyPermissionName{
			value: "can_delete",
		},
		CAN_DEPLOY: EnvironmentPermissionV2BodyPermissionName{
			value: "can_deploy",
		},
		CAN_MANAGE: EnvironmentPermissionV2BodyPermissionName{
			value: "can_manage",
		},
	}
}

func (c EnvironmentPermissionV2BodyPermissionName) Value() string {
	return c.value
}

func (c EnvironmentPermissionV2BodyPermissionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvironmentPermissionV2BodyPermissionName) UnmarshalJSON(b []byte) error {
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
