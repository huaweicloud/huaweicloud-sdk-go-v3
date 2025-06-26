package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerBudgetRecordsResponse Response Object
type ListSubCustomerBudgetRecordsResponse struct {

	// |参数名称：返回总条数。| |参数的约束及描述：范围限制：0-2147483647|
	TotalCount *int32 `json:"total_count,omitempty"`

	// |参数名称：金额单位。||参数的约束及描述：范围限制：0-3。1：元|
	MeasureId *int32 `json:"measure_id,omitempty"`

	// |参数名称：币种。||参数的约束及描述：范围限制：0-10。CNY：人民币，USD：美元|
	Currency *string `json:"currency,omitempty"`

	// |参数名称：客户预算设置记录列表。||参数的约束及描述：客户预算设置记录列表|
	RecordList     *[]BudgetRecordInfo `json:"record_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSubCustomerBudgetRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerBudgetRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomerBudgetRecordsResponse", string(data)}, " ")
}
