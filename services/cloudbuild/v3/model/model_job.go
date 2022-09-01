package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Job struct {

	// 任务ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 任务名称
	JobName *string `json:"job_name,omitempty" xml:"job_name"`

	// 任务创建者
	JobCreator *string `json:"job_creator,omitempty" xml:"job_creator"`

	// 用户名称
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 最新执行时间
	LastBuildTime float32 `json:"last_build_time,omitempty" xml:"last_build_time"`

	// 健康分值
	HealthScore *int32 `json:"health_score,omitempty" xml:"health_score"`

	// 代码来源
	SourceCode *string `json:"source_code,omitempty" xml:"source_code"`

	// 最新构建状态
	LastBuildStatus *string `json:"last_build_status,omitempty" xml:"last_build_status"`

	// 是否已结束
	IsFinished *bool `json:"is_finished,omitempty" xml:"is_finished"`

	// 是否已禁用
	Disabled *bool `json:"disabled,omitempty" xml:"disabled"`

	// 是否已收藏
	Favorite *bool `json:"favorite,omitempty" xml:"favorite"`

	// 是否有修改任务权限
	IsModify *bool `json:"is_modify,omitempty" xml:"is_modify"`

	// 是否有删除任务权限
	IsDelete *bool `json:"is_delete,omitempty" xml:"is_delete"`

	// 是否有执行任务权限
	IsExecute *bool `json:"is_execute,omitempty" xml:"is_execute"`

	// 是否有复制任务权限
	IsCopy *bool `json:"is_copy,omitempty" xml:"is_copy"`

	// 是否有禁用任务权限
	IsForbidden *bool `json:"is_forbidden,omitempty" xml:"is_forbidden"`

	// 是否有查看任务权限
	IsView *bool `json:"is_view,omitempty" xml:"is_view"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
