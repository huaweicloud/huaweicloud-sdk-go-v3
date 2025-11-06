package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDefaultReviewCategoriesResponse Response Object
type ListDefaultReviewCategoriesResponse struct {

	// **参数解释：** 检视意见分类(所有可勾选的，需传参with_default_review_categories: true才返回)。
	CodehubDefaultCategories *[]CategoryDto `json:"codehub_default_categories,omitempty"`

	// **参数解释：** 系统预置检视意见分类(需传参with_default_review_categories: true才返回)。
	HicodeDefaultCategories *[]CategoryDto `json:"hicode_default_categories,omitempty"`
	HttpStatusCode          int            `json:"-"`
}

func (o ListDefaultReviewCategoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDefaultReviewCategoriesResponse struct{}"
	}

	return strings.Join([]string{"ListDefaultReviewCategoriesResponse", string(data)}, " ")
}
