package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueListPiFilterInfo pi过滤条件
type IssueListPiFilterInfo struct {

	// 迭代列表
	Sprints *[]string `json:"sprints,omitempty"`

	// pi的id，层级关系：pi -> 迭代 -> 需求
	PiId *string `json:"pi_id,omitempty"`
}

func (o IssueListPiFilterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueListPiFilterInfo struct{}"
	}

	return strings.Join([]string{"IssueListPiFilterInfo", string(data)}, " ")
}
