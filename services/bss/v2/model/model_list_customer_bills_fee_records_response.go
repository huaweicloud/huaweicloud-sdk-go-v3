package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCustomerBillsFeeRecordsResponse struct {

	// 结果集数量，只有成功才返回这个参数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 资源费用记录数据。 具体请参见表2。
	Records *[]MonthlyBillRecord `json:"records,omitempty" xml:"records"`

	// 币种。 CNY：人民币
	Currency       *string `json:"currency,omitempty" xml:"currency"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCustomerBillsFeeRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerBillsFeeRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerBillsFeeRecordsResponse", string(data)}, " ")
}
