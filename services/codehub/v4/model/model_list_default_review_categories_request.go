package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDefaultReviewCategoriesRequest Request Object
type ListDefaultReviewCategoriesRequest struct {
}

func (o ListDefaultReviewCategoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDefaultReviewCategoriesRequest struct{}"
	}

	return strings.Join([]string{"ListDefaultReviewCategoriesRequest", string(data)}, " ")
}
