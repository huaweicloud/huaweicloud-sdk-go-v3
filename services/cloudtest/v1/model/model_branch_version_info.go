package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BranchVersionInfo struct {

	// 名称
	Name string `json:"name"`

	// 开发分支名称
	Number *string `json:"number,omitempty"`

	// 是否为Master分支
	IsMaster *string `json:"is_master,omitempty"`

	// PBI ID
	PbiId *string `json:"pbi_id,omitempty"`

	// PBI信息
	PbiName *string `json:"pbi_name,omitempty"`

	// 开始时间
	PlanStartDate *string `json:"plan_start_date,omitempty"`

	// 开始时间戳
	PlanStartTimestamp *int64 `json:"plan_start_timestamp,omitempty"`

	// 结束时间
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// 结束时间戳
	PlanEndTimestamp *int64 `json:"plan_end_timestamp,omitempty"`

	// 是否同步git库
	AsynGit *string `json:"asyn_git,omitempty"`

	// 项目ID（云龙场景，传入微服务ID）
	ProjectUuid string `json:"project_uuid"`

	// 项目名称（云龙场景，传入微服务名）
	ProjectName *string `json:"project_name,omitempty"`
}

func (o BranchVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BranchVersionInfo struct{}"
	}

	return strings.Join([]string{"BranchVersionInfo", string(data)}, " ")
}
