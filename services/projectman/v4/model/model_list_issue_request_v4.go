package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListIssueRequestV4 struct {

	// 处理人id
	AssignedIds *[]int32 `json:"assigned_ids,omitempty" xml:"assigned_ids"`

	// 创建者id
	CreatorIds *[]int32 `json:"creator_ids,omitempty" xml:"creator_ids"`

	// 开发人id,对应用户信息的数字id
	DeveloperIds *[]int32 `json:"developer_ids,omitempty" xml:"developer_ids"`

	// id, 领域, 14 '性能', 15 '功能', 16 '可靠性', 17 '网络安全', 18 '可维护性', 19 '其他DFX', 20 '可用性',
	DomainIds *[]int32 `json:"domain_ids,omitempty" xml:"domain_ids"`

	// 完成度
	DoneRatios *[]int32 `json:"done_ratios,omitempty" xml:"done_ratios"`

	// 迭代id
	IterationIds *[]int32 `json:"iteration_ids,omitempty" xml:"iteration_ids"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 分页索引，偏移量，offset是limit的整数倍数，limit=10,offset=0,10,20...
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 模块id
	ModuleIds *[]int32 `json:"module_ids,omitempty" xml:"module_ids"`

	// 优先级
	PriorityIds *[]int32 `json:"priority_ids,omitempty" xml:"priority_ids"`

	// 查询类型 backlog feature epic
	QueryType *string `json:"query_type,omitempty" xml:"query_type"`

	// 查询类型
	SeverityIds *[]int32 `json:"severity_ids,omitempty" xml:"severity_ids"`

	// 状态   id, 新建   1, 进行中 2, 已解决 3, 测试中 4, 已关闭 5, 已拒绝 6,
	StatusIds *[]int32 `json:"status_ids,omitempty" xml:"status_ids"`

	// 故事点id
	StoryPointIds *[]int32 `json:"story_point_ids,omitempty" xml:"story_point_ids"`

	// 工作项类型,2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
	TrackerIds *[]int32 `json:"tracker_ids,omitempty" xml:"tracker_ids"`

	// true 查询的工作项包含已经逻辑删除的，false 查询的工作项不包含已经删除的
	IncludeDeleted *bool `json:"include_deleted,omitempty" xml:"include_deleted"`

	// 根据工作项的更新时间查询工作项，(查询的起始时间,查询的结束时间)
	UpdatedTimeInterval *string `json:"updated_time_interval,omitempty" xml:"updated_time_interval"`

	// 自定义字段
	CustomFields *[]ListIssueRequestV4CustomFields `json:"custom_fields,omitempty" xml:"custom_fields"`
}

func (o ListIssueRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueRequestV4 struct{}"
	}

	return strings.Join([]string{"ListIssueRequestV4", string(data)}, " ")
}
