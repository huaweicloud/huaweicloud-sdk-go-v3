package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestCasesQueryInfo struct {

	// 关键字查询，用例名或编号
	Keyword *string `json:"keyword,omitempty"`

	// 执行平台
	Exeplatforms *[]string `json:"exeplatforms,omitempty"`

	// 是否是我的
	Own *bool `json:"own,omitempty"`

	UseOffset *bool `json:"useOffset,omitempty"`

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 用例URI集合
	CaseUris *[]string `json:"case_uris,omitempty"`

	// 处理者ID集合
	OwnerIds *[]string `json:"owner_ids,omitempty"`

	// 状态Code集合
	StatusCodes *[]string `json:"status_codes,omitempty"`

	// 用例等级ID集合
	RankIds *[]string `json:"rank_ids,omitempty"`

	// 模块ID集合
	ModuleIds *[]string `json:"module_ids,omitempty"`

	// 需求编号
	IssueId *string `json:"issue_id,omitempty"`

	// 创建者ID集合
	CreatorIds *[]string `json:"creator_ids,omitempty"`

	// 结果Code集合
	ResultCodes *[]string `json:"result_codes,omitempty"`

	// 归属迭代ID集合
	IterationIds *[]string `json:"iteration_ids,omitempty"`

	// 创建开始时间
	CreateStartTime *string `json:"create_start_time,omitempty"`

	// 创建结束时间
	CreateEndTime *string `json:"create_end_time,omitempty"`

	// 是否关联需求（null：不限，false：未关联，true：已关联）
	AssociatedIssue *bool `json:"associated_issue,omitempty"`

	// 是否关联缺陷（null：不限，false：未关联，true：已关联）
	AssociatedDefects *bool `json:"associated_defects,omitempty"`

	// 是否查询子需求关联的用例，默认true
	IncludeSubIssue *bool `json:"include_sub_issue,omitempty"`

	// 是否查询子目录的用例，默认true
	IncludeSubFeature *bool `json:"include_sub_feature,omitempty"`

	// 标签ID集合
	LabelIds *[]string `json:"label_ids,omitempty"`

	// 执行开始时间
	ExecuteStartTime *string `json:"execute_start_time,omitempty"`

	// 执行结束时间
	ExecuteEndTime *string `json:"execute_end_time,omitempty"`

	// 执行者ID集合
	ExecutorIds *[]string `json:"executor_ids,omitempty"`

	// 类型
	TestTypes *[]string `json:"test_types,omitempty"`

	// 是否组合关键字
	IsKeyword *bool `json:"is_keyword,omitempty"`

	// 是否是需求树点击的查询关联用例
	IssueTreeSearch *bool `json:"issue_tree_search,omitempty"`

	// 服务类型
	ServiceType *int32 `json:"service_type,omitempty"`

	// 服务类型集合
	ServiceTypes *[]int32 `json:"service_types,omitempty"`

	// 阶段过程（2：测试设计，3：测试执行，4：质量报告）
	StageType *int32 `json:"stage_type,omitempty"`

	// 目录URI
	FeatureUri *string `json:"feature_uri,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序方式
	SortType *string `json:"sort_type,omitempty"`

	// 当前页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 用例类型
	CaseType *int32 `json:"case_type,omitempty"`

	// 用例自定义字段信息
	CustomFieldInfo *[]QueryCustomFieldsInfo `json:"custom_field_info,omitempty"`

	// 测试套uri
	TaskUri *string `json:"task_uri,omitempty"`

	// 是否返回需求具体信息（返回需求名称，需求id）
	AssociateIssueDetail *bool `json:"associate_issue_detail,omitempty"`

	// 该字段为false,则查询全量用例，为true表示查询未分配测试套的用例
	NotAssignTask *bool `json:"not_assign_task,omitempty"`

	// 是否来自测试设计（null或者[true, false]：不限，[true]：来自测试设计，[false]：否来自测试设计）
	TestDesigns *[]bool `json:"test_designs,omitempty"`

	// 用例评审状态
	ReviewStatus *int32 `json:"review_status,omitempty"`
}

func (o TestCasesQueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCasesQueryInfo struct{}"
	}

	return strings.Join([]string{"TestCasesQueryInfo", string(data)}, " ")
}
