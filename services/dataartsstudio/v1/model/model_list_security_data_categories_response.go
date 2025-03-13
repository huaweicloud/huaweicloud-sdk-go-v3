package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDataCategoriesResponse Response Object
type ListSecurityDataCategoriesResponse struct {

	// 数据分类总的数量
	Total *int32 `json:"total,omitempty"`

	// 数据分类列表
	CategoryGroups *[]DataCategoryDto `json:"category_groups,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListSecurityDataCategoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDataCategoriesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityDataCategoriesResponse", string(data)}, " ")
}
