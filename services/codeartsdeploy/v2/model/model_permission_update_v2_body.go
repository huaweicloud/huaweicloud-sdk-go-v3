package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PermissionUpdateV2Body 主机相关修改权限请求体
type PermissionUpdateV2Body struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 角色id
	RoleId string `json:"role_id"`

	// 权限名称，can_view：查看权限；can_edit：编辑权限；can_delete：删除权限；can_add_host：添加主机权限；can_manage：权限管理权限；can_copy：复制主机权限
	PermissionName PermissionUpdateV2BodyPermissionName `json:"permission_name"`

	// true 有权限，false 无权限
	PermissionValue bool `json:"permission_value"`
}

func (o PermissionUpdateV2Body) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionUpdateV2Body struct{}"
	}

	return strings.Join([]string{"PermissionUpdateV2Body", string(data)}, " ")
}

type PermissionUpdateV2BodyPermissionName struct {
	value string
}

type PermissionUpdateV2BodyPermissionNameEnum struct {
	CAN_VIEW     PermissionUpdateV2BodyPermissionName
	CAN_EDIT     PermissionUpdateV2BodyPermissionName
	CAN_DELETE   PermissionUpdateV2BodyPermissionName
	CAN_ADD_HOST PermissionUpdateV2BodyPermissionName
	CAN_MANAGE   PermissionUpdateV2BodyPermissionName
	CAN_COPY     PermissionUpdateV2BodyPermissionName
}

func GetPermissionUpdateV2BodyPermissionNameEnum() PermissionUpdateV2BodyPermissionNameEnum {
	return PermissionUpdateV2BodyPermissionNameEnum{
		CAN_VIEW: PermissionUpdateV2BodyPermissionName{
			value: "can_view",
		},
		CAN_EDIT: PermissionUpdateV2BodyPermissionName{
			value: "can_edit",
		},
		CAN_DELETE: PermissionUpdateV2BodyPermissionName{
			value: "can_delete",
		},
		CAN_ADD_HOST: PermissionUpdateV2BodyPermissionName{
			value: "can_add_host",
		},
		CAN_MANAGE: PermissionUpdateV2BodyPermissionName{
			value: "can_manage",
		},
		CAN_COPY: PermissionUpdateV2BodyPermissionName{
			value: "can_copy",
		},
	}
}

func (c PermissionUpdateV2BodyPermissionName) Value() string {
	return c.value
}

func (c PermissionUpdateV2BodyPermissionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionUpdateV2BodyPermissionName) UnmarshalJSON(b []byte) error {
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
