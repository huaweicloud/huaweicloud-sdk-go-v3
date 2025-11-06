package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostNoteRequiredAttributeDto 检视意见设置请求体
type PostNoteRequiredAttributeDto struct {

	// **参数解释：** 是否勾选指派给。
	IsAssigneeIdRequired *bool `json:"is_assignee_id_required,omitempty"`

	// **参数解释：** 是否勾选意见分类。
	IsReviewCategoriesRequired *bool `json:"is_review_categories_required,omitempty"`

	// **参数解释：** 是否勾选意见模块。
	IsReviewModulesRequired *bool `json:"is_review_modules_required,omitempty"`
}

func (o PostNoteRequiredAttributeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostNoteRequiredAttributeDto struct{}"
	}

	return strings.Join([]string{"PostNoteRequiredAttributeDto", string(data)}, " ")
}
