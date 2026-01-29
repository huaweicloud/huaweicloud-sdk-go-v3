package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PrivilegeParam struct {

	// **参数解释**: 角色id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	RoleId string `json:"role_id"`

	// **参数解释**: 项目ID，可以从调用API处获取，也可以从控制台获取。获取方式请参考[获取项目ID](CloudArtifact_api_0015.xml)。 **约束限制**: 只能由英文字母、数字组成，且长度为32个字符。 **取值范围**: 不涉及。 **默认取值**: 无。
	ProjectId string `json:"project_id"`

	// **参数解释**: 地域服务id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	AreaServiceId string `json:"area_service_id"`

	// **参数解释**: 授权对象路径。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	GrantedObjectPath string `json:"granted_object_path"`

	// **参数解释**: 授权对象id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	GrantedObjectTypeId string `json:"granted_object_type_id"`

	// **参数解释**: 操作权限，多个权限以英文逗号隔开。 **约束限制**: 不涉及。 **取值范围**: 可选值：createrepository,editrepository,restore,deleterepository,physicdelete,restoreall,clearall,deleteorredeploy,downloadorview,import,upload,export。 **默认取值**: 无。
	Operations PrivilegeParamOperations `json:"operations"`
}

func (o PrivilegeParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivilegeParam struct{}"
	}

	return strings.Join([]string{"PrivilegeParam", string(data)}, " ")
}

type PrivilegeParamOperations struct {
	value string
}

type PrivilegeParamOperationsEnum struct {
	CREATEREPOSITORY PrivilegeParamOperations
	EDITREPOSITORY   PrivilegeParamOperations
	RESTORE          PrivilegeParamOperations
	DELETEREPOSITORY PrivilegeParamOperations
	PHYSICDELETE     PrivilegeParamOperations
	RESTOREALL       PrivilegeParamOperations
	CLEARALL         PrivilegeParamOperations
	DELETEORREDEPLOY PrivilegeParamOperations
	DOWNLOADORVIEW   PrivilegeParamOperations
	IMPORT           PrivilegeParamOperations
	UPLOAD           PrivilegeParamOperations
	EXPORT           PrivilegeParamOperations
}

func GetPrivilegeParamOperationsEnum() PrivilegeParamOperationsEnum {
	return PrivilegeParamOperationsEnum{
		CREATEREPOSITORY: PrivilegeParamOperations{
			value: "createrepository",
		},
		EDITREPOSITORY: PrivilegeParamOperations{
			value: "editrepository",
		},
		RESTORE: PrivilegeParamOperations{
			value: "restore",
		},
		DELETEREPOSITORY: PrivilegeParamOperations{
			value: "deleterepository",
		},
		PHYSICDELETE: PrivilegeParamOperations{
			value: "physicdelete",
		},
		RESTOREALL: PrivilegeParamOperations{
			value: "restoreall",
		},
		CLEARALL: PrivilegeParamOperations{
			value: "clearall",
		},
		DELETEORREDEPLOY: PrivilegeParamOperations{
			value: "deleteorredeploy",
		},
		DOWNLOADORVIEW: PrivilegeParamOperations{
			value: "downloadorview",
		},
		IMPORT: PrivilegeParamOperations{
			value: "import",
		},
		UPLOAD: PrivilegeParamOperations{
			value: "upload",
		},
		EXPORT: PrivilegeParamOperations{
			value: "export",
		},
	}
}

func (c PrivilegeParamOperations) Value() string {
	return c.value
}

func (c PrivilegeParamOperations) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivilegeParamOperations) UnmarshalJSON(b []byte) error {
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
