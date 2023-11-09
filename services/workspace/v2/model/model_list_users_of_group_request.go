package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListUsersOfGroupRequest Request Object
type ListUsersOfGroupRequest struct {

	// 用户名支持模糊查询。
	UserName *string `json:"user_name,omitempty"`

	// 用户组ID。
	GroupId string `json:"group_id"`

	// 用户描述支持模糊查询。
	Description *string `json:"description,omitempty"`

	// 激活类型。 - USER_ACTIVATE：用户激活 - ADMIN_ACTIVATE：管理员激活
	ActiveType *ListUsersOfGroupRequestActiveType `json:"active_type,omitempty"`

	// 用于分页查询，返回桌面数量限制。如果不指定或为0，默认2000，最大2000。
	Limit *string `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *string `json:"offset,omitempty"`
}

func (o ListUsersOfGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersOfGroupRequest struct{}"
	}

	return strings.Join([]string{"ListUsersOfGroupRequest", string(data)}, " ")
}

type ListUsersOfGroupRequestActiveType struct {
	value string
}

type ListUsersOfGroupRequestActiveTypeEnum struct {
	USER_ACTIVATE  ListUsersOfGroupRequestActiveType
	ADMIN_ACTIVATE ListUsersOfGroupRequestActiveType
}

func GetListUsersOfGroupRequestActiveTypeEnum() ListUsersOfGroupRequestActiveTypeEnum {
	return ListUsersOfGroupRequestActiveTypeEnum{
		USER_ACTIVATE: ListUsersOfGroupRequestActiveType{
			value: "USER_ACTIVATE",
		},
		ADMIN_ACTIVATE: ListUsersOfGroupRequestActiveType{
			value: "ADMIN_ACTIVATE",
		},
	}
}

func (c ListUsersOfGroupRequestActiveType) Value() string {
	return c.value
}

func (c ListUsersOfGroupRequestActiveType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUsersOfGroupRequestActiveType) UnmarshalJSON(b []byte) error {
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
