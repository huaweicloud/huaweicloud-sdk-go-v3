package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LimitInfoV2 struct {

	// 使用限制ID，主键。
	UseLimitiInfoId *string `json:"use_limiti_info_id,omitempty"`

	// 折扣限制，key的取值请参考表5。
	LimitKey *string `json:"limit_key,omitempty"`

	// value1。
	Value1 *string `json:"value1,omitempty"`

	// value2。
	Value2 *string `json:"value2,omitempty"`

	// value单位。
	ValueUnit *string `json:"value_unit,omitempty"`

	// 限制类型。
	LimitType *string `json:"limit_type,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`
}

func (o LimitInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LimitInfoV2 struct{}"
	}

	return strings.Join([]string{"LimitInfoV2", string(data)}, " ")
}
