package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// DevUcEnvironmentPermission 应用环境鉴权类
type DevUcEnvironmentPermission struct {

	// 权限id
	Id *int32 `json:"id,omitempty"`

	// 角色id
	RoleId *string `json:"role_id,omitempty"`

	// 角色id列表
	DevucRoleIdList *[]string `json:"devuc_role_id_list,omitempty"`

	// 角色类型， environment-creator： 环境创建者； project： 项目管理员；template-customized-inst：系统角色； template-project-customized、project-customized：自定义角色
	RoleType *DevUcEnvironmentPermissionRoleType `json:"role_type,omitempty"`

	// 角色名称
	Name *string `json:"name,omitempty"`

	// 局点信息
	Region *string `json:"region,omitempty"`

	// 环境id
	EnvironmentId *string `json:"environment_id,omitempty"`

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`

	// 是否有编辑权限
	CanEdit *bool `json:"can_edit,omitempty"`

	// 是否有删除权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有部署权限
	CanDeploy *bool `json:"can_deploy,omitempty"`

	// 是否有权限管理权限
	CanManage *bool `json:"can_manage,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o DevUcEnvironmentPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevUcEnvironmentPermission struct{}"
	}

	return strings.Join([]string{"DevUcEnvironmentPermission", string(data)}, " ")
}

type DevUcEnvironmentPermissionRoleType struct {
	value string
}

type DevUcEnvironmentPermissionRoleTypeEnum struct {
	ENVIRONMENT_CREATOR         DevUcEnvironmentPermissionRoleType
	PROJECT                     DevUcEnvironmentPermissionRoleType
	TEMPLATE_CUSTOMIZED_INST    DevUcEnvironmentPermissionRoleType
	TEMPLATE_PROJECT_CUSTOMIZED DevUcEnvironmentPermissionRoleType
	PROJECT_CUSTOMIZED          DevUcEnvironmentPermissionRoleType
}

func GetDevUcEnvironmentPermissionRoleTypeEnum() DevUcEnvironmentPermissionRoleTypeEnum {
	return DevUcEnvironmentPermissionRoleTypeEnum{
		ENVIRONMENT_CREATOR: DevUcEnvironmentPermissionRoleType{
			value: "environment-creator",
		},
		PROJECT: DevUcEnvironmentPermissionRoleType{
			value: "project",
		},
		TEMPLATE_CUSTOMIZED_INST: DevUcEnvironmentPermissionRoleType{
			value: "template-customized-inst",
		},
		TEMPLATE_PROJECT_CUSTOMIZED: DevUcEnvironmentPermissionRoleType{
			value: "template-project-customized",
		},
		PROJECT_CUSTOMIZED: DevUcEnvironmentPermissionRoleType{
			value: "project-customized",
		},
	}
}

func (c DevUcEnvironmentPermissionRoleType) Value() string {
	return c.value
}

func (c DevUcEnvironmentPermissionRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DevUcEnvironmentPermissionRoleType) UnmarshalJSON(b []byte) error {
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
