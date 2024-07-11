package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// DevUcClusterPermission 主机集群权限类
type DevUcClusterPermission struct {

	// 局点信息
	Region *string `json:"region,omitempty"`

	// 角色id
	RoleId *string `json:"role_id,omitempty"`

	// 角色id列表
	DevucRoleIdList *[]string `json:"devuc_role_id_list,omitempty"`

	// 角色名称
	Name *string `json:"name,omitempty"`

	// 主机集群id
	GroupId *string `json:"group_id,omitempty"`

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`

	// 是否有编辑权限
	CanEdit *bool `json:"can_edit,omitempty"`

	// 是否有删除权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有添加主机权限
	CanAddHost *bool `json:"can_add_host,omitempty"`

	// 是否有权限管理权限
	CanManage *bool `json:"can_manage,omitempty"`

	// 是否有拷贝权限
	CanCopy *bool `json:"can_copy,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 角色类型，project-customized：自定义角色；template-project-customized：系统自定义角色； template-customized-inst：系统角色；cluster-creator：集群创建者；project_admin：项目创建者
	RoleType *DevUcClusterPermissionRoleType `json:"role_type,omitempty"`
}

func (o DevUcClusterPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevUcClusterPermission struct{}"
	}

	return strings.Join([]string{"DevUcClusterPermission", string(data)}, " ")
}

type DevUcClusterPermissionRoleType struct {
	value string
}

type DevUcClusterPermissionRoleTypeEnum struct {
	PROJECT_CUSTOMIZED          DevUcClusterPermissionRoleType
	TEMPLATE_PROJECT_CUSTOMIZED DevUcClusterPermissionRoleType
	TEMPLATE_CUSTOMIZED_INST    DevUcClusterPermissionRoleType
	CLUSTER_CREATOR             DevUcClusterPermissionRoleType
	PROJECT_ADMIN               DevUcClusterPermissionRoleType
}

func GetDevUcClusterPermissionRoleTypeEnum() DevUcClusterPermissionRoleTypeEnum {
	return DevUcClusterPermissionRoleTypeEnum{
		PROJECT_CUSTOMIZED: DevUcClusterPermissionRoleType{
			value: "project-customized",
		},
		TEMPLATE_PROJECT_CUSTOMIZED: DevUcClusterPermissionRoleType{
			value: "template-project-customized",
		},
		TEMPLATE_CUSTOMIZED_INST: DevUcClusterPermissionRoleType{
			value: "template-customized-inst",
		},
		CLUSTER_CREATOR: DevUcClusterPermissionRoleType{
			value: "cluster-creator",
		},
		PROJECT_ADMIN: DevUcClusterPermissionRoleType{
			value: "project_admin",
		},
	}
}

func (c DevUcClusterPermissionRoleType) Value() string {
	return c.value
}

func (c DevUcClusterPermissionRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DevUcClusterPermissionRoleType) UnmarshalJSON(b []byte) error {
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
