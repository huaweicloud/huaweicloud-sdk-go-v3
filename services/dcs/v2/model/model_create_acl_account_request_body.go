package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAclAccountRequestBody 创建ACL账号
type CreateAclAccountRequestBody struct {

	// 账号名称 - 以字母开头。 - 内容由字母、数字、中划线、下划线组成。 - 长度范围[1~64]个字符。
	AccountName string `json:"account_name"`

	// 账号权限
	AccountRole CreateAclAccountRequestBodyAccountRole `json:"account_role"`

	// 账号密码 - 输入长度为8到64位的字符串。 - 不能包含正序逆序用户名。 - 必须包含如下四种字符中的三种组合（不允许包含冒号）：   - 小写字母   - 大写字母   - 数字   - 特殊字符包括（`~!@#$%^&*()-_=+\\|{},<.>/?）
	AccountPassword string `json:"account_password"`

	// 账号描述
	Description *string `json:"description,omitempty"`
}

func (o CreateAclAccountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAclAccountRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAclAccountRequestBody", string(data)}, " ")
}

type CreateAclAccountRequestBodyAccountRole struct {
	value string
}

type CreateAclAccountRequestBodyAccountRoleEnum struct {
	READ  CreateAclAccountRequestBodyAccountRole
	WRITE CreateAclAccountRequestBodyAccountRole
}

func GetCreateAclAccountRequestBodyAccountRoleEnum() CreateAclAccountRequestBodyAccountRoleEnum {
	return CreateAclAccountRequestBodyAccountRoleEnum{
		READ: CreateAclAccountRequestBodyAccountRole{
			value: "read",
		},
		WRITE: CreateAclAccountRequestBodyAccountRole{
			value: "write",
		},
	}
}

func (c CreateAclAccountRequestBodyAccountRole) Value() string {
	return c.value
}

func (c CreateAclAccountRequestBodyAccountRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAclAccountRequestBodyAccountRole) UnmarshalJSON(b []byte) error {
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
