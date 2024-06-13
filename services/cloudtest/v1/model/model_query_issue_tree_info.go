package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryIssueTreeInfo struct {

	// 服务类型
	ServiceType *int32 `json:"service_type,omitempty"`

	// 服务类型集合
	ServiceTypes *[]int32 `json:"service_types,omitempty"`

	// 父节点id
	ParentId *string `json:"parent_id,omitempty"`

	// 页码
	PageNumber *int32 `json:"page_number,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`

	Filter *IssueListFilterInfo `json:"filter,omitempty"`

	// trackerId
	TrackerId *string `json:"tracker_id,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	// 任务udi
	TaskUri *string `json:"task_uri,omitempty"`

	// 是否统计子需求的用例数，默认true
	IncludeSubIssue *bool `json:"include_sub_issue,omitempty"`
}

func (o QueryIssueTreeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryIssueTreeInfo struct{}"
	}

	return strings.Join([]string{"QueryIssueTreeInfo", string(data)}, " ")
}
