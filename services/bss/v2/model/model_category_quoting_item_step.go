package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

// CategoryQuotingItemStep 分类报价项阶梯信息
type CategoryQuotingItemStep struct {

	// 阶梯ID
	StepId *string `json:"step_id,omitempty"`

	// 阶梯编号
	StepNo *string `json:"step_no,omitempty"`

	// 阶梯起始值
	StepStart *decimal.Decimal `json:"step_start,omitempty"`

	// 起始值度量单位
	StartMeasureId *int32 `json:"start_measure_id,omitempty"`

	// 阶梯结束值
	StepEnd *decimal.Decimal `json:"step_end,omitempty"`

	// 结束值度量单位
	EndMeasureId *int32 `json:"end_measure_id,omitempty"`

	// 阶梯生效时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 阶梯失效时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	ExpireTime *string `json:"expire_time,omitempty"`

	// 运营站点编码
	SiteCode *string `json:"site_code,omitempty"`
}

func (o CategoryQuotingItemStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryQuotingItemStep struct{}"
	}

	return strings.Join([]string{"CategoryQuotingItemStep", string(data)}, " ")
}
