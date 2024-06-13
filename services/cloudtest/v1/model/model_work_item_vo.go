package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkItemVo struct {

	// 工作项名称
	Name *string `json:"name,omitempty"`

	// 处理人
	Owner *string `json:"owner,omitempty"`

	// 工作项路径
	Path *string `json:"path,omitempty"`

	// 预计开始日期
	StartDate *string `json:"start_date,omitempty"`

	// 预计结束日期
	DueDate *string `json:"due_date,omitempty"`

	// 逻辑region，外部使用公有云实际区域，内部使用默认值
	Region *string `json:"region,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 更新人
	Updator *string `json:"updator,omitempty"`

	// 项目ID，外部使用项目ID，内部使用默认值
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 工作项编号
	WorkItemId *string `json:"work_item_id,omitempty"`

	// 状态ID
	StatusId *string `json:"status_id,omitempty"`

	// 状态
	StatusName *string `json:"status_name,omitempty"`

	// 类型ID
	TrackerId *string `json:"tracker_id,omitempty"`

	// 类型
	TrackerName *string `json:"tracker_name,omitempty"`

	// 迭代ID
	IterationId *string `json:"iteration_id,omitempty"`

	// 模块ID
	ModuleId *string `json:"module_id,omitempty"`

	// 重要程度ID
	SeverityId *string `json:"severity_id,omitempty"`

	// 重要程度
	SeverityName *string `json:"severity_name,omitempty"`

	// 父工作项编号
	ParentWorkitemId *string `json:"parent_workitem_id,omitempty"`

	// 看板ID
	BoardId *string `json:"board_id,omitempty"`

	// 看板
	BoardName *string `json:"board_name,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 迭代名
	IterationName *string `json:"iteration_name,omitempty"`

	// 模块名
	ModuleName *string `json:"module_name,omitempty"`

	// 模块path
	ModulePath *string `json:"module_path,omitempty"`

	// 模块路径名称
	ModulePathName *string `json:"module_path_name,omitempty"`

	// 处理人
	OwnerName *string `json:"owner_name,omitempty"`

	// 父工作项下是否有子工作项包含动态
	HaveChildDynamic *bool `json:"have_child_dynamic,omitempty"`

	// 父工作项下是否有子工作项
	HasChild *bool `json:"has_child,omitempty"`

	// 需求动态数量
	IssueDynamicCount *int32 `json:"issue_dynamic_count,omitempty"`

	// 用例数量
	CaseCount *int32 `json:"case_count,omitempty"`

	// xBoard项目工作项序列号
	SequenceId *string `json:"sequence_id,omitempty"`

	// pi的id，层级关系：pi -> 迭代 -> 需求
	PiId *string `json:"pi_id,omitempty"`

	// 迭代ID
	PiName *string `json:"pi_name,omitempty"`
}

func (o WorkItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemVo struct{}"
	}

	return strings.Join([]string{"WorkItemVo", string(data)}, " ")
}
