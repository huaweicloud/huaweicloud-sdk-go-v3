package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMergeRequestDiscussionBodyDto struct {

	// 检视意见内容
	Body string `json:"body"`

	// 严重程度
	Severity *string `json:"severity,omitempty"`

	// 指派人id
	AssigneeId *string `json:"assignee_id,omitempty"`

	// 检视意见分类
	ReviewCategories *string `json:"review_categories,omitempty"`

	// 检视意见模块
	ReviewModules *string `json:"review_modules,omitempty"`

	// 提出人id
	ProposerId *string `json:"proposer_id,omitempty"`

	Position *PositionDto `json:"position,omitempty"`
}

func (o CreateMergeRequestDiscussionBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestDiscussionBodyDto struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestDiscussionBodyDto", string(data)}, " ")
}
