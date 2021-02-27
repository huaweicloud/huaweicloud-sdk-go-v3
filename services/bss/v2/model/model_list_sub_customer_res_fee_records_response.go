package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSubCustomerResFeeRecordsResponse struct {
	// 客户的消费记录数据。 具体请参见表2。
	FeeRecords *[]SubCustomerResFeeRecordV2 `json:"fee_records,omitempty"`
	// 结果集数量，只有成功才返回这个参数。
	Count *int32 `json:"count,omitempty"`
	// 币种。 CNY：人民币
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSubCustomerResFeeRecordsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubCustomerResFeeRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomerResFeeRecordsResponse", string(data)}, " ")
}
