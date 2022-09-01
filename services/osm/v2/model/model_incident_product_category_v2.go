package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentProductCategoryV2 struct {

	// 产品类型id
	IncidentProductCategoryId *string `json:"incident_product_category_id,omitempty" xml:"incident_product_category_id"`

	// 产品类型名称
	IncidentProductCategoryName *string `json:"incident_product_category_name,omitempty" xml:"incident_product_category_name"`

	// 产品类型描述
	IncidentProductCategoryDesc *string `json:"incident_product_category_desc,omitempty" xml:"incident_product_category_desc"`

	// 产品类型简称
	IncidentProductCategoryAcronym *string `json:"incident_product_category_acronym,omitempty" xml:"incident_product_category_acronym"`

	// 是否可以使用支持计划权益
	CanUseSupportPlan *bool `json:"can_use_support_plan,omitempty" xml:"can_use_support_plan"`
}

func (o IncidentProductCategoryV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentProductCategoryV2 struct{}"
	}

	return strings.Join([]string{"IncidentProductCategoryV2", string(data)}, " ")
}
