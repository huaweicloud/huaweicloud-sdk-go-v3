package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueListFilterInfo 过滤条件
type IssueListFilterInfo struct {

	// 迭代id列表
	IterationIds *[]string `json:"iteration_ids,omitempty"`

	// pi过滤条件
	PiSprints *[]IssueListPiFilterInfo `json:"pi_sprints,omitempty"`

	// 需求名
	Subject *string `json:"subject,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	// 需求状态id
	StatusId *string `json:"status_id,omitempty"`
}

func (o IssueListFilterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueListFilterInfo struct{}"
	}

	return strings.Join([]string{"IssueListFilterInfo", string(data)}, " ")
}
