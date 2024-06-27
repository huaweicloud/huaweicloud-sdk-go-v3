package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueTreeInfo 查询迭代的需求树Body
type IssueTreeInfo struct {

	// 过滤条件：处理人
	Owner *string `json:"owner,omitempty"`

	// 需求ID
	IssueId *string `json:"issue_id,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页展示条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 过滤条件：迭代ID
	IterationId *string `json:"iteration_id,omitempty"`

	// 过滤条件：重要程度ID
	SeverityId *string `json:"severity_id,omitempty"`

	// 过滤条件：状态ID
	StatusId *string `json:"status_id,omitempty"`

	// 过滤条件：模块ID
	ModuleId *string `json:"module_id,omitempty"`

	// 过滤条件：状态ID多个条件，每个条件取或，-2代表搜索未设置
	StatusIds *string `json:"status_ids,omitempty"`

	// 过滤条件：模块ID多个，每个条件取或，-2代表搜索未设置
	ModuleIds *string `json:"module_ids,omitempty"`

	// 迭代、pi过滤条件
	PiFilter *[]IssueListPiFilterInfo `json:"pi_filter,omitempty"`

	// 状态名称列表
	StatusNames *[]string `json:"status_names,omitempty"`
}

func (o IssueTreeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueTreeInfo struct{}"
	}

	return strings.Join([]string{"IssueTreeInfo", string(data)}, " ")
}
