package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RunAuthorizationActionRequestBody struct {

	// 被赋权的用户名称，该用户将有权访问指定的DLI资源权限，被收回或者更新访问权限。
	UserName *string `json:"user_name,omitempty"`

	// 被赋权的项目ID，数据赋权给其他项目后，该项目的管理员将 有权访问指定的DLI资源权限，被收回或者更新访问权限。
	GrantProjectId *string `json:"grant_project_id,omitempty"`

	// 指定赋权或回收。值为：grant，revoke或update。  说明：当用户同时拥有grant和revoke权限的时候才有权限使用update操作。
	Action RunAuthorizationActionRequestBodyAction `json:"action"`

	// 赋权信息。具体参数请参考Privilege参数。
	Privileges []Privilege `json:"privileges"`
}

func (o RunAuthorizationActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAuthorizationActionRequestBody struct{}"
	}

	return strings.Join([]string{"RunAuthorizationActionRequestBody", string(data)}, " ")
}

type RunAuthorizationActionRequestBodyAction struct {
	value string
}

type RunAuthorizationActionRequestBodyActionEnum struct {
	GRANT  RunAuthorizationActionRequestBodyAction
	REVOKE RunAuthorizationActionRequestBodyAction
	UPDATE RunAuthorizationActionRequestBodyAction
}

func GetRunAuthorizationActionRequestBodyActionEnum() RunAuthorizationActionRequestBodyActionEnum {
	return RunAuthorizationActionRequestBodyActionEnum{
		GRANT: RunAuthorizationActionRequestBodyAction{
			value: "grant",
		},
		REVOKE: RunAuthorizationActionRequestBodyAction{
			value: "revoke",
		},
		UPDATE: RunAuthorizationActionRequestBodyAction{
			value: "update",
		},
	}
}

func (c RunAuthorizationActionRequestBodyAction) Value() string {
	return c.value
}

func (c RunAuthorizationActionRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunAuthorizationActionRequestBodyAction) UnmarshalJSON(b []byte) error {
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
