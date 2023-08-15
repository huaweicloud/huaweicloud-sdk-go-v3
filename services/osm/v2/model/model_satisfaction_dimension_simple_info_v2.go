package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SatisfactionDimensionSimpleInfoV2 struct {

	// 总的分数
	Value *int32 `json:"value,omitempty"`

	// 满意度id
	SatisfactionId *int32 `json:"satisfaction_id,omitempty"`

	// 满意度名称
	SatisfactionName *string `json:"satisfaction_name,omitempty"`

	// 满意度描述
	SatisfactionDesc *string `json:"satisfaction_desc,omitempty"`

	// 每格的分数
	PerValue *int32 `json:"per_value,omitempty"`

	// 满意度分类id
	SatCategoryId *string `json:"sat_category_id,omitempty"`
}

func (o SatisfactionDimensionSimpleInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SatisfactionDimensionSimpleInfoV2 struct{}"
	}

	return strings.Join([]string{"SatisfactionDimensionSimpleInfoV2", string(data)}, " ")
}
