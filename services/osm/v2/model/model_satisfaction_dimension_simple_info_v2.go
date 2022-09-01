package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SatisfactionDimensionSimpleInfoV2 struct {

	// 总的分数
	Value *int32 `json:"value,omitempty" xml:"value"`

	// 满意度id
	SatisfactionId *int32 `json:"satisfaction_id,omitempty" xml:"satisfaction_id"`

	// 满意度名称
	SatisfactionName *string `json:"satisfaction_name,omitempty" xml:"satisfaction_name"`

	// 满意度描述
	SatisfactionDesc *string `json:"satisfaction_desc,omitempty" xml:"satisfaction_desc"`

	// 每格的分数
	PerValue *int32 `json:"per_value,omitempty" xml:"per_value"`

	// 满意度分类id
	SatCategoryId *string `json:"sat_category_id,omitempty" xml:"sat_category_id"`
}

func (o SatisfactionDimensionSimpleInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SatisfactionDimensionSimpleInfoV2 struct{}"
	}

	return strings.Join([]string{"SatisfactionDimensionSimpleInfoV2", string(data)}, " ")
}
