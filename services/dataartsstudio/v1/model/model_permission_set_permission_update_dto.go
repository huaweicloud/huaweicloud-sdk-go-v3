package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionSetPermissionUpdateDto struct {

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 权限操作列表
	PermissionActions *[]PermissionSetPermissionUpdateDtoPermissionActions `json:"permission_actions,omitempty"`
}

func (o PermissionSetPermissionUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSetPermissionUpdateDto struct{}"
	}

	return strings.Join([]string{"PermissionSetPermissionUpdateDto", string(data)}, " ")
}

type PermissionSetPermissionUpdateDtoPermissionActions struct {
	value string
}

type PermissionSetPermissionUpdateDtoPermissionActionsEnum struct {
	ALL    PermissionSetPermissionUpdateDtoPermissionActions
	SELECT PermissionSetPermissionUpdateDtoPermissionActions
	UPDATE PermissionSetPermissionUpdateDtoPermissionActions
	CREATE PermissionSetPermissionUpdateDtoPermissionActions
	DROP   PermissionSetPermissionUpdateDtoPermissionActions
	ALTER  PermissionSetPermissionUpdateDtoPermissionActions
	INDEX  PermissionSetPermissionUpdateDtoPermissionActions
	LOCK   PermissionSetPermissionUpdateDtoPermissionActions
	READ   PermissionSetPermissionUpdateDtoPermissionActions
	WRITE  PermissionSetPermissionUpdateDtoPermissionActions
}

func GetPermissionSetPermissionUpdateDtoPermissionActionsEnum() PermissionSetPermissionUpdateDtoPermissionActionsEnum {
	return PermissionSetPermissionUpdateDtoPermissionActionsEnum{
		ALL: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "ALL",
		},
		SELECT: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "SELECT",
		},
		UPDATE: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "UPDATE",
		},
		CREATE: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "CREATE",
		},
		DROP: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "DROP",
		},
		ALTER: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "ALTER",
		},
		INDEX: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "INDEX",
		},
		LOCK: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "LOCK",
		},
		READ: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "READ",
		},
		WRITE: PermissionSetPermissionUpdateDtoPermissionActions{
			value: "WRITE",
		},
	}
}

func (c PermissionSetPermissionUpdateDtoPermissionActions) Value() string {
	return c.value
}

func (c PermissionSetPermissionUpdateDtoPermissionActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionSetPermissionUpdateDtoPermissionActions) UnmarshalJSON(b []byte) error {
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
