package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerCouponChangeRecordsResponse Response Object
type ListCustomerCouponChangeRecordsResponse struct {

	// |参数名称：返回总条数| |参数的约束及描述：返回总条数|
	TotalCount *int32 `json:"total_count,omitempty"`

	// |参数名称：币种| |参数约束及描述：CNY：人民币。|
	Currency *string `json:"currency,omitempty"`

	// |参数名称：优惠券收支明细记录列表| |参数约束以及描述：优惠券收支明细记录列表，具体请参见表 CustomerCouponChangeRecord。|
	Records        *[]CustomerCouponChangeRecord `json:"records,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListCustomerCouponChangeRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerCouponChangeRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerCouponChangeRecordsResponse", string(data)}, " ")
}
