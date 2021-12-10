package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RoleAction struct {
	// 动作名

	ActionCname *string `json:"action_cname,omitempty"`
	// 执行动作 。 - CREATE_INSTANCE 创建实例 - DELETE_INSTANCE 删除实例 - UPDATE_INSTANCE 修改实例 - QUERY_INSTANCE 查询实例列表 - RUN_INSTANCE 运行实例 - CREATE_SUB_ORG 创建子组织 - UPDATE_SUB_ORG 修改子组织 - DELETE_SUB_ORG 删除子组织 - QUERY_SUB_ORG 查询子组织列表 - ADD_USER_TO_ORG 新增组织中用户 - SET_USER_ROLES 设置用户角色 - DELETE_USER_IN_ORG 删除组织中用户 - QUERY_USER_IN_ORG 查询组织中用户列表

	Actions *RoleActionActions `json:"actions,omitempty"`
	// id

	Id *string `json:"id,omitempty"`
	// 角色id

	RoleId *string `json:"role_id,omitempty"`
}

func (o RoleAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleAction struct{}"
	}

	return strings.Join([]string{"RoleAction", string(data)}, " ")
}

type RoleActionActions struct {
	value string
}

type RoleActionActionsEnum struct {
	CREATE_INSTANCE    RoleActionActions
	DELETE_INSTANCE    RoleActionActions
	UPDATE_INSTANCE    RoleActionActions
	QUERY_INSTANCE     RoleActionActions
	RUN_INSTANCE       RoleActionActions
	CREATE_SUB_ORG     RoleActionActions
	UPDATE_SUB_ORG     RoleActionActions
	DELETE_SUB_ORG     RoleActionActions
	QUERY_SUB_ORG      RoleActionActions
	ADD_USER_TO_ORG    RoleActionActions
	SET_USER_ROLES     RoleActionActions
	DELETE_USER_IN_ORG RoleActionActions
	QUERY_USER_IN_ORG  RoleActionActions
}

func GetRoleActionActionsEnum() RoleActionActionsEnum {
	return RoleActionActionsEnum{
		CREATE_INSTANCE: RoleActionActions{
			value: "CREATE_INSTANCE",
		},
		DELETE_INSTANCE: RoleActionActions{
			value: "DELETE_INSTANCE",
		},
		UPDATE_INSTANCE: RoleActionActions{
			value: "UPDATE_INSTANCE",
		},
		QUERY_INSTANCE: RoleActionActions{
			value: "QUERY_INSTANCE",
		},
		RUN_INSTANCE: RoleActionActions{
			value: "RUN_INSTANCE",
		},
		CREATE_SUB_ORG: RoleActionActions{
			value: "CREATE_SUB_ORG",
		},
		UPDATE_SUB_ORG: RoleActionActions{
			value: "UPDATE_SUB_ORG",
		},
		DELETE_SUB_ORG: RoleActionActions{
			value: "DELETE_SUB_ORG",
		},
		QUERY_SUB_ORG: RoleActionActions{
			value: "QUERY_SUB_ORG",
		},
		ADD_USER_TO_ORG: RoleActionActions{
			value: "ADD_USER_TO_ORG",
		},
		SET_USER_ROLES: RoleActionActions{
			value: "SET_USER_ROLES",
		},
		DELETE_USER_IN_ORG: RoleActionActions{
			value: "DELETE_USER_IN_ORG",
		},
		QUERY_USER_IN_ORG: RoleActionActions{
			value: "QUERY_USER_IN_ORG",
		},
	}
}

func (c RoleActionActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RoleActionActions) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
