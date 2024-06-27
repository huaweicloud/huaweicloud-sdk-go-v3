package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTaskAssignCasesInfo struct {

	// 测试用例在任务中的阶段
	Stage *int32 `json:"stage,omitempty"`

	// 处理人过滤数组
	Owners *[]string `json:"owners,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 页数量
	PageSize *int32 `json:"page_size,omitempty"`

	// 结果过滤
	Results *[]string `json:"results,omitempty"`

	// 状态过滤
	Status *[]string `json:"status,omitempty"`

	// 分支/迭代uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 任务版本过滤条件，影响关联任务的结果查询，查询当前任务版本下的用例最新结果
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序方法
	SortType *string `json:"sort_type,omitempty"`

	// 特性目录URI
	FeatureUri *string `json:"feature_uri,omitempty"`

	// 测试套结果uri
	TaskResultUri *string `json:"task_result_uri,omitempty"`

	// 用例等级ID集合
	RankIds *[]int32 `json:"rank_ids,omitempty"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 需求id
	IssueId *string `json:"issue_id,omitempty"`

	// 是否关联需求（null：不限，false：未关联，true：已关联）
	AssociatedIssue *bool `json:"associated_issue,omitempty"`

	// 是否全选所有页（null：不全选，false：不全选，true：全选），用于任务批量执行结果功能，只返回用例uri，不返回其他信息
	SelectAllPages *bool `json:"select_all_pages,omitempty"`

	// 用例是否可用
	IsAvailable *bool `json:"is_available,omitempty"`

	// 用例脚本字段是否有值
	IsScriptExist *bool `json:"is_script_exist,omitempty"`
}

func (o QueryTaskAssignCasesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTaskAssignCasesInfo struct{}"
	}

	return strings.Join([]string{"QueryTaskAssignCasesInfo", string(data)}, " ")
}
