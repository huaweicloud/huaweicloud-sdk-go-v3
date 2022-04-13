package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ShowDeployTaskDetailResponse struct {
	// 部署任务id

	TaskId *string `json:"task_id,omitempty"`
	// 部署任务名称

	Name *string `json:"name,omitempty"`
	// devcloud项目id

	ProjectId *string `json:"project_id,omitempty"`
	// devcloud项目名称

	ProjectName *string `json:"project_name,omitempty"`
	// 部署类型模式，包括deployTemplate，ansible，shell

	DeploySystem *string `json:"deploy_system,omitempty"`
	// 创建时间

	CreateTime *string `json:"create_time,omitempty"`
	// 修改时间

	UpdateTime *string `json:"update_time,omitempty"`
	// 部署任务状态，Draft表示草稿状态，Available表示可用状态

	State *ShowDeployTaskDetailResponseState `json:"state,omitempty"`
	// 最后一次执行时间

	ExecutionTime *string `json:"execution_time,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 是否使用默认权限矩阵

	IsDefautPermission *bool `json:"is_defaut_permission,omitempty"`
	// 模板id

	TemplateId *string `json:"template_id,omitempty"`
	// 部署任务创建者用户名

	Owner *string `json:"owner,omitempty"`
	// 部署任务创建者昵称

	NickName *string `json:"nick_name,omitempty"`
	// 部署任务创建者用户ID

	OwnerId *string `json:"owner_id,omitempty"`
	// 部署任务创建者租户ID

	TenantId *string `json:"tenant_id,omitempty"`
	// 部署任务创建者租户名

	TenantName *string `json:"tenant_name,omitempty"`
	// slave集群id，默认为null时使用devcloud八爪鱼slave集群，用户自定义slave时为slave集群id

	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`
	// 当前用户是否已收藏

	IsCare *bool `json:"is_care,omitempty"`
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
	// 部署任务和应用组件对应关系

	AppComponentList *[]AppComponentDao `json:"app_component_list,omitempty"`
	// 角色ID,0：部署任务创建者，-1：项目创建者，3：项目经理，4：开发人员，5：测试经理，6：测试人员，7：参与者，8：浏览者

	RoleId         *int32 `json:"role_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDeployTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeployTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDeployTaskDetailResponse", string(data)}, " ")
}

type ShowDeployTaskDetailResponseState struct {
	value string
}

type ShowDeployTaskDetailResponseStateEnum struct {
	AVAILABLE ShowDeployTaskDetailResponseState
	DRAFT     ShowDeployTaskDetailResponseState
}

func GetShowDeployTaskDetailResponseStateEnum() ShowDeployTaskDetailResponseStateEnum {
	return ShowDeployTaskDetailResponseStateEnum{
		AVAILABLE: ShowDeployTaskDetailResponseState{
			value: "Available",
		},
		DRAFT: ShowDeployTaskDetailResponseState{
			value: "Draft",
		},
	}
}

func (c ShowDeployTaskDetailResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDeployTaskDetailResponseState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
