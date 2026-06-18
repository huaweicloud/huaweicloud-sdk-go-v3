package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReviewSettingResponse Response Object
type CreateReviewSettingResponse struct {

	// **参数解释：** 是否启用启用检视意见分类与模块。
	CategoriesAndModulesEnabled *bool `json:"categories_and_modules_enabled,omitempty"`

	// **参数解释：** 是否启用系统预置检视意见分类。
	SecondaryCategoryEnabled *bool `json:"secondary_category_enabled,omitempty"`

	// **参数解释：** 检视意见分类(已勾选)。
	PrimaryCategories *[]CategoryDto `json:"primary_categories,omitempty"`

	// **参数解释：** 检视意见分类的key(已勾选)。
	ReviewDefaultCategories *[]string `json:"review_default_categories,omitempty"`

	// **参数解释：** 自定义分类。
	ReviewCustomizedCategories *[]string `json:"review_customized_categories,omitempty"`

	// **参数解释：** 检视意见模块。
	ReviewModules *[]string `json:"review_modules,omitempty"`

	// **参数解释：** 系统预置检视意见分类类型(启用系统预置检视意见分类时返回，默认'HiCode')。
	SecondaryCategoryType *string `json:"secondary_category_type,omitempty"`

	// **参数解释：** 系统预置检视意见分类详情(启用系统预置检视意见分类时返回)。
	SecondaryCategories *[]CategoryDto `json:"secondary_categories,omitempty"`

	// **参数解释：** 检视意见必填项(post接口不更新必填项时不返回)。
	NoteRequiredAttributes *[]RequiredAttributeDto `json:"note_required_attributes,omitempty"`

	// **参数解释：** 检视意见分类(可勾选的，需get接口传参with_default_review_categories: true才返回)。
	CodehubDefaultCategories *[]CategoryDto `json:"codehub_default_categories,omitempty"`

	// **参数解释：** 系统预置检视意见分类(需get接口传参with_default_review_categories: true才返回)。
	HicodeDefaultCategories *[]CategoryDto `json:"hicode_default_categories,omitempty"`
	HttpStatusCode          int            `json:"-"`
}

func (o CreateReviewSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReviewSettingResponse struct{}"
	}

	return strings.Join([]string{"CreateReviewSettingResponse", string(data)}, " ")
}
