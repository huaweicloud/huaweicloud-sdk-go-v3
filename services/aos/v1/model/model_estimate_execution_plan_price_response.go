package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateExecutionPlanPriceResponse Response Object
type EstimateExecutionPlanPriceResponse struct {

	// 币种，枚举值   * [`CNY` - 元，中国站返回的币种](tag:hws)[`USD` - 美元，国际站返回的币种](tag:hws_hk)[`USD` - 美元，欧洲站返回的币种](tag:hws_eu)
	Currency *string `json:"currency,omitempty"`

	// 执行计划中所有资源的询价结果
	Items          *[]ItemsResponse `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o EstimateExecutionPlanPriceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateExecutionPlanPriceResponse struct{}"
	}

	return strings.Join([]string{"EstimateExecutionPlanPriceResponse", string(data)}, " ")
}
