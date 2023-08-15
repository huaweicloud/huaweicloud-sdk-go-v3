package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIssueInfoVo 关联需求信息
type AssociateIssueInfoVo struct {

	// 是否已关联
	Associate *bool `json:"associate,omitempty"`

	// 需求ID
	IssueId *string `json:"issue_id,omitempty"`

	// 需求类型
	TrackerId *string `json:"tracker_id,omitempty"`

	// 工作项层级ID
	BoardId *string `json:"board_id,omitempty"`

	// 需求类型名称
	TrackerName *string `json:"tracker_name,omitempty"`
}

func (o AssociateIssueInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIssueInfoVo struct{}"
	}

	return strings.Join([]string{"AssociateIssueInfoVo", string(data)}, " ")
}
