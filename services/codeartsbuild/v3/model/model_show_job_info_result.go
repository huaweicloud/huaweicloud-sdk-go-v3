package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobInfoResult 任务信息
type ShowJobInfoResult struct {

	// 构建任务ID
	Id *string `json:"id,omitempty"`

	// 构建工程ID,唯一对应codeci_job_id
	ProjectId *string `json:"project_id,omitempty"`

	// 使用项目权限
	ProjectPermissionSwitch *bool `json:"project_permission_switch,omitempty"`

	// 执行时间
	BuildTime *string `json:"build_time,omitempty"`

	// 收费时间
	ChargeTime *string `json:"charge_time,omitempty"`

	// 任务创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 是否已禁用
	Disabled *bool `json:"disabled,omitempty"`

	// 是否已收藏
	Favorite *bool `json:"favorite,omitempty"`

	// 代码来源
	SourceCode *string `json:"source_code,omitempty"`

	// 运行状态
	RunningStatus *string `json:"running_status,omitempty"`

	// 重新运行
	NewBuild *bool `json:"new_build,omitempty"`

	// 任务名称
	JobName *string `json:"job_name,omitempty"`

	// 是否有复制任务权限
	IsCopy *bool `json:"is_copy,omitempty"`

	// 是否有删除任务权限
	IsDelete *bool `json:"is_delete,omitempty"`

	// 是否有执行任务权限
	IsExecute *bool `json:"is_execute,omitempty"`

	// 是否有禁用任务权限
	IsForbidden *bool `json:"is_forbidden,omitempty"`

	// 是否有管理任务权限
	IsManager *bool `json:"is_manager,omitempty"`

	// 是否有修改任务权限
	IsModify *bool `json:"is_modify,omitempty"`

	// 是否有查看任务权限
	IsView *bool `json:"is_view,omitempty"`

	// 最终构建状态
	LastBuildStatus *string `json:"last_build_status,omitempty"`

	// 最后构建时间
	LastBuildTime *int64 `json:"last_build_time,omitempty"`

	// 健康度
	HealthScore *int32 `json:"health_score,omitempty"`
}

func (o ShowJobInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInfoResult struct{}"
	}

	return strings.Join([]string{"ShowJobInfoResult", string(data)}, " ")
}
