package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplicationPermissionVo 角色应用权限
type ApplicationPermissionVo struct {

	// 是否有编辑权限
	CanModify *bool `json:"can_modify,omitempty"`

	// 是否有删除的权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`

	// 是否有执行权限
	CanExecute *bool `json:"can_execute,omitempty"`

	// 是否有复制权限
	CanCopy *bool `json:"can_copy,omitempty"`

	// 是否有管理权限，包含增删改查执行以及权限修改
	CanManage *bool `json:"can_manage,omitempty"`

	// 是否有新建环境权限
	CanCreateEnv *bool `json:"can_create_env,omitempty"`

	// 是否有禁用权限
	CanDisable *bool `json:"can_disable,omitempty"`

	// 角色名称
	Name *string `json:"name,omitempty"`

	// 局点信息
	Region *string `json:"region,omitempty"`

	// 角色id
	RoleId *string `json:"role_id,omitempty"`

	// 角色类型， app-creator： 应用创建者； project： 项目管理员；template-customized-inst：系统角色； template-project-customized、project-customized：自定义角色
	RoleType *ApplicationPermissionVoRoleType `json:"role_type,omitempty"`
}

func (o ApplicationPermissionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationPermissionVo struct{}"
	}

	return strings.Join([]string{"ApplicationPermissionVo", string(data)}, " ")
}

type ApplicationPermissionVoRoleType struct {
	value string
}

type ApplicationPermissionVoRoleTypeEnum struct {
	APP_CREATOR                 ApplicationPermissionVoRoleType
	PROJECT                     ApplicationPermissionVoRoleType
	TEMPLATE_CUSTOMIZED_INST    ApplicationPermissionVoRoleType
	TEMPLATE_PROJECT_CUSTOMIZED ApplicationPermissionVoRoleType
	PROJECT_CUSTOMIZED          ApplicationPermissionVoRoleType
}

func GetApplicationPermissionVoRoleTypeEnum() ApplicationPermissionVoRoleTypeEnum {
	return ApplicationPermissionVoRoleTypeEnum{
		APP_CREATOR: ApplicationPermissionVoRoleType{
			value: "app-creator",
		},
		PROJECT: ApplicationPermissionVoRoleType{
			value: "project",
		},
		TEMPLATE_CUSTOMIZED_INST: ApplicationPermissionVoRoleType{
			value: "template-customized-inst",
		},
		TEMPLATE_PROJECT_CUSTOMIZED: ApplicationPermissionVoRoleType{
			value: "template-project-customized",
		},
		PROJECT_CUSTOMIZED: ApplicationPermissionVoRoleType{
			value: "project-customized",
		},
	}
}

func (c ApplicationPermissionVoRoleType) Value() string {
	return c.value
}

func (c ApplicationPermissionVoRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplicationPermissionVoRoleType) UnmarshalJSON(b []byte) error {
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
