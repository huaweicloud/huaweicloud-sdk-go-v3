package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCostsResponse Response Object
type ListCostsResponse struct {

	// 货币。 USD：美元
	Currency *string `json:"currency,omitempty"`

	// 总条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 按天或按月的明细金额。
	CostData       *[]CostDataByDimension `json:"cost_data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListCostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCostsResponse struct{}"
	}

	return strings.Join([]string{"ListCostsResponse", string(data)}, " ")
}
