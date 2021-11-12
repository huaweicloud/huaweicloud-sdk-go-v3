package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetSubCustomerDiscountV2 struct {
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId string `json:"customer_id"`
	// 折扣率，最高精确到2位小数。 折扣范围：0.8~1。 如果折扣率是85%，则折扣率写成0.85。  说明： 折扣为1表示不打折，相当于删除伙伴折扣。

	Discount float64 `json:"discount"`
	// 生效时间。仅discount=1时无需填写。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。

	EffectiveTime *string `json:"effective_time,omitempty"`
	// 失效时间。仅discount=1时无需填写。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。

	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o SetSubCustomerDiscountV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSubCustomerDiscountV2 struct{}"
	}

	return strings.Join([]string{"SetSubCustomerDiscountV2", string(data)}, " ")
}
