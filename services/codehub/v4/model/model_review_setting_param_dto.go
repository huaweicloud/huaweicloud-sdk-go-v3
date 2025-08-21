package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReviewSettingParamDto 检视意见设置请求体
type ReviewSettingParamDto struct {

	// **参数解释：** 是否启用启用检视意见分类与模块。
	CategoriesAndModulesEnabled *bool `json:"categories_and_modules_enabled,omitempty"`

	// **参数解释：** 检视意见模块。
	ReviewModules *[]string `json:"review_modules,omitempty"`

	// **参数解释：** 是否启用系统预置检视意见分类。
	SecondaryCategoryEnabled *bool `json:"secondary_category_enabled,omitempty"`

	// **参数解释：** 检视意见分类(已勾选的分类的key列表)。
	ReviewDefaultCategories *[]string `json:"review_default_categories,omitempty"`

	// **参数解释：** 自定义分类列表。
	ReviewCustomizedCategories *[]string `json:"review_customized_categories,omitempty"`

	// **参数解释：** 是否勾选指派给。
	IsAssigneeIdRequired *bool `json:"is_assignee_id_required,omitempty"`

	// **参数解释：** 是否勾选意见分类。
	IsReviewCategoriesRequired *bool `json:"is_review_categories_required,omitempty"`

	// **参数解释：** 是否勾选意见模块。
	IsReviewModulesRequired *bool `json:"is_review_modules_required,omitempty"`
}

func (o ReviewSettingParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewSettingParamDto struct{}"
	}

	return strings.Join([]string{"ReviewSettingParamDto", string(data)}, " ")
}
