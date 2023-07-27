package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReviewSettingDto struct {

	// 是否开启检视意见分类和模块
	CategoriesAndModulesEnabled *bool `json:"categories_and_modules_enabled,omitempty"`

	// 是否开启二级分类
	SecondaryCategoryEnabled *bool `json:"secondary_category_enabled,omitempty"`

	// 是否禁止关联issue
	ForbiddenAddToIssue *bool `json:"forbidden_add_to_issue,omitempty"`

	// 一级分类
	PrimaryCategories *[]CategoryDto `json:"primary_categories,omitempty"`

	// 默认分类
	ReviewDefaultCategories *[]string `json:"review_default_categories,omitempty"`

	// 自定义分类
	ReviewCustomizedCategories *[]string `json:"review_customized_categories,omitempty"`

	// 检视意见模块
	ReviewModules *[]string `json:"review_modules,omitempty"`

	// 目标id
	SourceId *int32 `json:"source_id,omitempty"`

	// 目标类型
	SourceType *string `json:"source_type,omitempty"`

	// 目标路径
	SourcePath *string `json:"source_path,omitempty"`

	// 二级分类类型
	SecondaryCategoryType *string `json:"secondary_category_type,omitempty"`

	// 二级分类
	SecondaryCategories *[]CategoryDto `json:"secondary_categories,omitempty"`
}

func (o ReviewSettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewSettingDto struct{}"
	}

	return strings.Join([]string{"ReviewSettingDto", string(data)}, " ")
}
