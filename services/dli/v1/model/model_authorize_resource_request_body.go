package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuthorizeResourceRequestBody struct {

	// 被赋权的用户名称，该用户将有权访问指定的DLI资源权限，被收回或者更新访问权限。
	UserName *string `json:"user_name,omitempty"`

	// 被赋权的项目ID，数据赋权给其他项目后，该项目的管理员将 有权访问指定的DLI资源权限，被收回或者更新访问权限。
	ProjectId *string `json:"projectId,omitempty"`

	// 指定赋权或回收。值为：grant，revoke或update。  说明：当用户同时拥有grant和revoke权限的时候才有权限使用update操作。
	Action AuthorizeResourceRequestBodyAction `json:"action"`

	// 赋权信息。具体参数请参考Privilege参数。
	Privileges []Privilege `json:"privileges"`
}

func (o AuthorizeResourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeResourceRequestBody struct{}"
	}

	return strings.Join([]string{"AuthorizeResourceRequestBody", string(data)}, " ")
}

type AuthorizeResourceRequestBodyAction struct {
	value string
}

type AuthorizeResourceRequestBodyActionEnum struct {
	GRANT  AuthorizeResourceRequestBodyAction
	REVOKE AuthorizeResourceRequestBodyAction
	UPDATE AuthorizeResourceRequestBodyAction
}

func GetAuthorizeResourceRequestBodyActionEnum() AuthorizeResourceRequestBodyActionEnum {
	return AuthorizeResourceRequestBodyActionEnum{
		GRANT: AuthorizeResourceRequestBodyAction{
			value: "grant",
		},
		REVOKE: AuthorizeResourceRequestBodyAction{
			value: "revoke",
		},
		UPDATE: AuthorizeResourceRequestBodyAction{
			value: "update",
		},
	}
}

func (c AuthorizeResourceRequestBodyAction) Value() string {
	return c.value
}

func (c AuthorizeResourceRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeResourceRequestBodyAction) UnmarshalJSON(b []byte) error {
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
