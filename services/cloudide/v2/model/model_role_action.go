package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RoleAction struct {
	// 动作名

	ActionCname *string `json:"action_cname,omitempty"`
	// 执行动作

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
