package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AclAccountRoleModifyBody struct {

	// 账号权限
	AccountRole *AclAccountRoleModifyBodyAccountRole `json:"account_role,omitempty"`
}

func (o AclAccountRoleModifyBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclAccountRoleModifyBody struct{}"
	}

	return strings.Join([]string{"AclAccountRoleModifyBody", string(data)}, " ")
}

type AclAccountRoleModifyBodyAccountRole struct {
	value string
}

type AclAccountRoleModifyBodyAccountRoleEnum struct {
	READ  AclAccountRoleModifyBodyAccountRole
	WRITE AclAccountRoleModifyBodyAccountRole
}

func GetAclAccountRoleModifyBodyAccountRoleEnum() AclAccountRoleModifyBodyAccountRoleEnum {
	return AclAccountRoleModifyBodyAccountRoleEnum{
		READ: AclAccountRoleModifyBodyAccountRole{
			value: "read",
		},
		WRITE: AclAccountRoleModifyBodyAccountRole{
			value: "write",
		},
	}
}

func (c AclAccountRoleModifyBodyAccountRole) Value() string {
	return c.value
}

func (c AclAccountRoleModifyBodyAccountRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AclAccountRoleModifyBodyAccountRole) UnmarshalJSON(b []byte) error {
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
