package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateHostGroupPermissionsResponse Response Object
type UpdateHostGroupPermissionsResponse struct {

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
	RoleType       *UpdateHostGroupPermissionsResponseRoleType `json:"role_type,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o UpdateHostGroupPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostGroupPermissionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostGroupPermissionsResponse", string(data)}, " ")
}

type UpdateHostGroupPermissionsResponseRoleType struct {
	value string
}

type UpdateHostGroupPermissionsResponseRoleTypeEnum struct {
	PROJECT_CUSTOMIZED          UpdateHostGroupPermissionsResponseRoleType
	TEMPLATE_PROJECT_CUSTOMIZED UpdateHostGroupPermissionsResponseRoleType
	TEMPLATE_CUSTOMIZED_INST    UpdateHostGroupPermissionsResponseRoleType
	CLUSTER_CREATOR             UpdateHostGroupPermissionsResponseRoleType
	PROJECT_ADMIN               UpdateHostGroupPermissionsResponseRoleType
}

func GetUpdateHostGroupPermissionsResponseRoleTypeEnum() UpdateHostGroupPermissionsResponseRoleTypeEnum {
	return UpdateHostGroupPermissionsResponseRoleTypeEnum{
		PROJECT_CUSTOMIZED: UpdateHostGroupPermissionsResponseRoleType{
			value: "project-customized",
		},
		TEMPLATE_PROJECT_CUSTOMIZED: UpdateHostGroupPermissionsResponseRoleType{
			value: "template-project-customized",
		},
		TEMPLATE_CUSTOMIZED_INST: UpdateHostGroupPermissionsResponseRoleType{
			value: "template-customized-inst",
		},
		CLUSTER_CREATOR: UpdateHostGroupPermissionsResponseRoleType{
			value: "cluster-creator",
		},
		PROJECT_ADMIN: UpdateHostGroupPermissionsResponseRoleType{
			value: "project_admin",
		},
	}
}

func (c UpdateHostGroupPermissionsResponseRoleType) Value() string {
	return c.value
}

func (c UpdateHostGroupPermissionsResponseRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHostGroupPermissionsResponseRoleType) UnmarshalJSON(b []byte) error {
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
