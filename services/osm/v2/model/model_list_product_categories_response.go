package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProductCategoriesResponse struct {
	// 总数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 产品类型列表

	IncidentProductCategoryList *[]IncidentProductCategoryV2 `json:"incident_product_category_list,omitempty"`
	HttpStatusCode              int                          `json:"-"`
}

func (o ListProductCategoriesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProductCategoriesResponse struct{}"
	}

	return strings.Join([]string{"ListProductCategoriesResponse", string(data)}, " ")
}
