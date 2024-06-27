package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IteratorVersionInfo struct {

	// 名称
	Name string `json:"name"`

	// 开发分支名称
	Number *string `json:"number,omitempty"`

	// 处理者ID
	Owner *string `json:"owner,omitempty"`

	// 待测版本
	Version *string `json:"version,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 关联迭代
	Iterations *[]string `json:"iterations,omitempty"`

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

	// 当前所处阶段
	CurrentStage *string `json:"current_stage,omitempty"`

	// 测试类型
	ServiceTypes *[]string `json:"service_types,omitempty"`

	// 关联需求
	IssueList *[]WorkItemInfo `json:"issue_list,omitempty"`

	// 风险等级
	RiskRating *int32 `json:"risk_rating,omitempty"`

	// 风险描述
	RiskDes *string `json:"risk_des,omitempty"`

	// 编辑风险信息标记
	IsUpdateRisk *string `json:"is_update_risk,omitempty"`

	// pi的id，层级关系：pi -> 迭代 -> 需求
	PiId *string `json:"pi_id,omitempty"`
}

func (o IteratorVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IteratorVersionInfo struct{}"
	}

	return strings.Join([]string{"IteratorVersionInfo", string(data)}, " ")
}
