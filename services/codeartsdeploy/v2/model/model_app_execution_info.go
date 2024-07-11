package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppExecutionInfo 查询应用详情响应体
type AppExecutionInfo struct {

	// 应用id
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 部署时间
	Duration *string `json:"duration,omitempty"`

	// 当前应用是否被禁用
	IsDisable *bool `json:"is_disable,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

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

	// 部署类型模式，包括deployTemplate、ansible、shell
	DeploySystem *string `json:"deploy_system,omitempty"`

	// 应用创建者用户id
	CreateUserId *string `json:"create_user_id,omitempty"`

	// 应用创建者租户id
	CreateTenantId *string `json:"create_tenant_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 最后一次部署时间
	ExecutionTime *string `json:"execution_time,omitempty"`

	// 部署结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 部署状态
	ExecutionState *string `json:"execution_state,omitempty"`

	// 部署记录序列号
	ReleaseId *int32 `json:"release_id,omitempty"`

	// 部署者id
	ExecutorId *string `json:"executor_id,omitempty"`

	// 部署者名称
	ExecutorNickName *string `json:"executor_nick_name,omitempty"`

	// 部署任务信息
	ArrangeInfos *[]TaskBaseResponseBody `json:"arrange_infos,omitempty"`
}

func (o AppExecutionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppExecutionInfo struct{}"
	}

	return strings.Join([]string{"AppExecutionInfo", string(data)}, " ")
}
