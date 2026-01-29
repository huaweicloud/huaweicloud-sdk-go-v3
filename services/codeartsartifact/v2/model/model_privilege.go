package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Privilege struct {

	// **参数解释**: 角色id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	RoleId string `json:"role_id"`

	// **参数解释**: 角色名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	RoleName *string `json:"role_name,omitempty"`

	// **参数解释**: 角色中文名。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	RoleChineseName *string `json:"role_chinese_name,omitempty"`

	// **参数解释**: 项目id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	ProjectId string `json:"project_id"`

	// **参数解释**: 地域服务id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	AreaServiceId string `json:"area_service_id"`

	// **参数解释**: 授权对象路径。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	GrantedObjectPath string `json:"granted_object_path"`

	// **参数解释**: 授权对象id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	GrantedObjectTypeId string `json:"granted_object_type_id"`

	// **参数解释**: 操作权限，多个权限以英文逗号隔开。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	Operations PrivilegeOperations `json:"operations"`

	// **参数解释**: 操作权限索引。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	OperationsIndex *[]int32 `json:"operations_index,omitempty"`
}

func (o Privilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Privilege struct{}"
	}

	return strings.Join([]string{"Privilege", string(data)}, " ")
}

type PrivilegeOperations struct {
	value string
}

type PrivilegeOperationsEnum struct {
	CREATEREPOSITORY PrivilegeOperations
	EDITREPOSITORY   PrivilegeOperations
	RESTORE          PrivilegeOperations
	DELETEREPOSITORY PrivilegeOperations
	PHYSICDELETE     PrivilegeOperations
	RESTOREALL       PrivilegeOperations
	CLEARALL         PrivilegeOperations
	DELETEORREDEPLOY PrivilegeOperations
	DOWNLOADORVIEW   PrivilegeOperations
	IMPORT           PrivilegeOperations
	UPLOAD           PrivilegeOperations
	EXPORT           PrivilegeOperations
}

func GetPrivilegeOperationsEnum() PrivilegeOperationsEnum {
	return PrivilegeOperationsEnum{
		CREATEREPOSITORY: PrivilegeOperations{
			value: "createrepository",
		},
		EDITREPOSITORY: PrivilegeOperations{
			value: "editrepository",
		},
		RESTORE: PrivilegeOperations{
			value: "restore",
		},
		DELETEREPOSITORY: PrivilegeOperations{
			value: "deleterepository",
		},
		PHYSICDELETE: PrivilegeOperations{
			value: "physicdelete",
		},
		RESTOREALL: PrivilegeOperations{
			value: "restoreall",
		},
		CLEARALL: PrivilegeOperations{
			value: "clearall",
		},
		DELETEORREDEPLOY: PrivilegeOperations{
			value: "deleteorredeploy",
		},
		DOWNLOADORVIEW: PrivilegeOperations{
			value: "downloadorview",
		},
		IMPORT: PrivilegeOperations{
			value: "import",
		},
		UPLOAD: PrivilegeOperations{
			value: "upload",
		},
		EXPORT: PrivilegeOperations{
			value: "export",
		},
	}
}

func (c PrivilegeOperations) Value() string {
	return c.value
}

func (c PrivilegeOperations) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivilegeOperations) UnmarshalJSON(b []byte) error {
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
