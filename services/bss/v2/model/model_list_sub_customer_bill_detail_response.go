package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerBillDetailResponse Response Object
type ListSubCustomerBillDetailResponse struct {

	// 结果集数量，只有成功才返回这个参数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 货币单位代码： CNY：人民币
	Currency *string `json:"currency,omitempty"`

	// 资源费用记录数据。 具体请参见表2。
	FeeRecords     *[]SubCustomerMonthlyBillDetail `json:"fee_records,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListSubCustomerBillDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerBillDetailResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomerBillDetailResponse", string(data)}, " ")
}
