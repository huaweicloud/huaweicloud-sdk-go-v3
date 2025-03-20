package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAndDeletePrivilegeReq struct {

	// 执行的操作(create|delete)
	Operation string `json:"operation"`

	// 权限标识
	Privilege *CreateAndDeletePrivilegeReqPrivilege `json:"privilege,omitempty"`
}

func (o CreateAndDeletePrivilegeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndDeletePrivilegeReq struct{}"
	}

	return strings.Join([]string{"CreateAndDeletePrivilegeReq", string(data)}, " ")
}

type CreateAndDeletePrivilegeReqPrivilege struct {
	value string
}

type CreateAndDeletePrivilegeReqPrivilegeEnum struct {
	EXPORT      CreateAndDeletePrivilegeReqPrivilege
	CREATE_CASE CreateAndDeletePrivilegeReqPrivilege
}

func GetCreateAndDeletePrivilegeReqPrivilegeEnum() CreateAndDeletePrivilegeReqPrivilegeEnum {
	return CreateAndDeletePrivilegeReqPrivilegeEnum{
		EXPORT: CreateAndDeletePrivilegeReqPrivilege{
			value: "export",
		},
		CREATE_CASE: CreateAndDeletePrivilegeReqPrivilege{
			value: "createCase",
		},
	}
}

func (c CreateAndDeletePrivilegeReqPrivilege) Value() string {
	return c.value
}

func (c CreateAndDeletePrivilegeReqPrivilege) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAndDeletePrivilegeReqPrivilege) UnmarshalJSON(b []byte) error {
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
