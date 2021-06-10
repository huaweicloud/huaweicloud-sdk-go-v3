package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCaseCategoriesResponse struct {
	// 总数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 工单子类型列表

	IncidentSubTypeList *[]IncidentSubTypeV2Do `json:"incident_sub_type_list,omitempty"`
	HttpStatusCode      int                    `json:"-"`
}

func (o ListCaseCategoriesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCaseCategoriesResponse struct{}"
	}

	return strings.Join([]string{"ListCaseCategoriesResponse", string(data)}, " ")
}
