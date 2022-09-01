package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentSatisfactionV2Do struct {

	// 满意度总分数
	Value *int32 `json:"value,omitempty" xml:"value"`

	// 满意度分类id
	SatisfactionId int32 `json:"satisfaction_id" xml:"satisfaction_id"`

	// 满意度的值
	SatisfactionValue int32 `json:"satisfaction_value" xml:"satisfaction_value"`

	// 满意度分类名称
	SatisfactionName *string `json:"satisfaction_name,omitempty" xml:"satisfaction_name"`

	// 每格的分数
	PerValue *int32 `json:"per_value,omitempty" xml:"per_value"`

	// 满意度维度id
	SatCategoryId *string `json:"sat_category_id,omitempty" xml:"sat_category_id"`

	// 满意度维度名称
	SatCategoryName *string `json:"sat_category_name,omitempty" xml:"sat_category_name"`
}

func (o IncidentSatisfactionV2Do) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentSatisfactionV2Do struct{}"
	}

	return strings.Join([]string{"IncidentSatisfactionV2Do", string(data)}, " ")
}
