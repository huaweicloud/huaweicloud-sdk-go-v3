package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskV2Info 部署任务详情信息
type TaskV2Info struct {

	// 部署任务id
	Id *string `json:"id,omitempty"`

	// 部署任务名称
	Name *string `json:"name,omitempty"`

	// 部署任务状态
	State *string `json:"state,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 部署任务所属人
	Owner *string `json:"owner,omitempty"`

	// 部署步骤
	Steps map[string]Step `json:"steps,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 部署类型模式，包括deployTemplate、ansible、shell
	DeploySystem *string `json:"deploy_system,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 角色id
	RoleId *int32 `json:"role_id,omitempty"`

	// 是否为默认角色
	IsDefautPermission *bool `json:"is_defaut_permission,omitempty"`

	// 模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 应用创建者昵称
	NickName *string `json:"nick_name,omitempty"`

	// 应用创建者用户id
	OwnerId *string `json:"owner_id,omitempty"`

	// 应用创建者租户id
	TenantId *string `json:"tenant_id,omitempty"`

	// 应用创建者租户名
	TenantName *string `json:"tenant_name,omitempty"`

	// slave集群id，默认为null时使用默认slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	// 当前用户是否已收藏
	IsCare *bool `json:"is_care,omitempty"`

	// 是否有编辑权限
	CanModify *bool `json:"can_modify,omitempty"`

	// 是否有删除的权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`

	// 是否有部署权限
	CanExecute *bool `json:"can_execute,omitempty"`

	// 是否有复制权限
	CanCopy *bool `json:"can_copy,omitempty"`

	// 是否有编辑应用权限矩阵的权限
	CanManage *bool `json:"can_manage,omitempty"`

	// 是否有创建环境的权限
	CanCreateEnv *bool `json:"can_create_env,omitempty"`

	// 是否有禁用应用的权限
	CanDisable *bool `json:"can_disable,omitempty"`

	// 应用组件列表
	AppComponentList *[]AppComponentDao `json:"app_component_list,omitempty"`

	// 部署记录序列号
	ReleaseId *int32 `json:"release_id,omitempty"`

	// 部署任务所属应用id
	AppId *string `json:"app_id,omitempty"`

	// 当前应用是否被禁用
	IsDisable *bool `json:"is_disable,omitempty"`
}

func (o TaskV2Info) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskV2Info struct{}"
	}

	return strings.Join([]string{"TaskV2Info", string(data)}, " ")
}
