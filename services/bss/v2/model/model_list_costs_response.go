package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCostsResponse struct {

	// 货币。 CNY：人民币
	Currency *string `json:"currency,omitempty" xml:"currency"`

	// 总条数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 按天或按月的明细金额。
	CostData       *[]CostDataByDimension `json:"cost_data,omitempty" xml:"cost_data"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListCostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCostsResponse struct{}"
	}

	return strings.Join([]string{"ListCostsResponse", string(data)}, " ")
}
