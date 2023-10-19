package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessRequest CheckPoliciesInput body
type AccessRequest struct {
	Resource *ResourceInput `json:"resource"`

	// 授权主体信息
	Principal []Principal `json:"principal"`

	// 权限信息,ALL,CREATE,ALTER,DROP,DESCRIBE,EXEC,CREATE_DATABASE,LIST_DATABASE,CREATE_TABLE,LIST_TABLE,CREATE_FUNC,LIST_FUNC,REGISTER_MODEL,LIST_MODEL,INSERT,UPDATE,DELETE,SELECT,READ,WRITE,OPERATE,USE
	Action AccessRequestAction `json:"action"`
}

func (o AccessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessRequest struct{}"
	}

	return strings.Join([]string{"AccessRequest", string(data)}, " ")
}

type AccessRequestAction struct {
	value string
}

type AccessRequestActionEnum struct {
	ALL             AccessRequestAction
	CREATE          AccessRequestAction
	ALTER           AccessRequestAction
	DROP            AccessRequestAction
	DESCRIBE        AccessRequestAction
	EXEC            AccessRequestAction
	CREATE_DATABASE AccessRequestAction
	LIST_DATABASE   AccessRequestAction
	CREATE_TABLE    AccessRequestAction
	LIST_TABLE      AccessRequestAction
	CREATE_FUNC     AccessRequestAction
	LIST_FUNC       AccessRequestAction
	REGISTER_MODEL  AccessRequestAction
	LIST_MODEL      AccessRequestAction
	INSERT          AccessRequestAction
	UPDATE          AccessRequestAction
	DELETE          AccessRequestAction
	SELECT          AccessRequestAction
	READ            AccessRequestAction
	WRITE           AccessRequestAction
	OPERATE         AccessRequestAction
	USE             AccessRequestAction
}

func GetAccessRequestActionEnum() AccessRequestActionEnum {
	return AccessRequestActionEnum{
		ALL: AccessRequestAction{
			value: "ALL",
		},
		CREATE: AccessRequestAction{
			value: "CREATE",
		},
		ALTER: AccessRequestAction{
			value: "ALTER",
		},
		DROP: AccessRequestAction{
			value: "DROP",
		},
		DESCRIBE: AccessRequestAction{
			value: "DESCRIBE",
		},
		EXEC: AccessRequestAction{
			value: "EXEC",
		},
		CREATE_DATABASE: AccessRequestAction{
			value: "CREATE_DATABASE",
		},
		LIST_DATABASE: AccessRequestAction{
			value: "LIST_DATABASE",
		},
		CREATE_TABLE: AccessRequestAction{
			value: "CREATE_TABLE",
		},
		LIST_TABLE: AccessRequestAction{
			value: "LIST_TABLE",
		},
		CREATE_FUNC: AccessRequestAction{
			value: "CREATE_FUNC",
		},
		LIST_FUNC: AccessRequestAction{
			value: "LIST_FUNC",
		},
		REGISTER_MODEL: AccessRequestAction{
			value: "REGISTER_MODEL",
		},
		LIST_MODEL: AccessRequestAction{
			value: "LIST_MODEL",
		},
		INSERT: AccessRequestAction{
			value: "INSERT",
		},
		UPDATE: AccessRequestAction{
			value: "UPDATE",
		},
		DELETE: AccessRequestAction{
			value: "DELETE",
		},
		SELECT: AccessRequestAction{
			value: "SELECT",
		},
		READ: AccessRequestAction{
			value: "READ",
		},
		WRITE: AccessRequestAction{
			value: "WRITE",
		},
		OPERATE: AccessRequestAction{
			value: "OPERATE",
		},
		USE: AccessRequestAction{
			value: "USE",
		},
	}
}

func (c AccessRequestAction) Value() string {
	return c.value
}

func (c AccessRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessRequestAction) UnmarshalJSON(b []byte) error {
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
