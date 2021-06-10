package model

import (
	"encoding/json"

	"strings"
)

type IncidentSubTypeV2Do struct {
	// 工单子类型id

	IncidentSubTypeId *string `json:"incident_sub_type_id,omitempty"`
	// 工单子类型名称

	IncidentSubTypeName *string `json:"incident_sub_type_name,omitempty"`
	// 产品类型列表

	IncidentProductCategoryList *[]IncidentProductCategoryV2 `json:"incident_product_category_list,omitempty"`
}

func (o IncidentSubTypeV2Do) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IncidentSubTypeV2Do struct{}"
	}

	return strings.Join([]string{"IncidentSubTypeV2Do", string(data)}, " ")
}
