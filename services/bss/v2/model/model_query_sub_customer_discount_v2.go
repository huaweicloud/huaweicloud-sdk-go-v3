package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuerySubCustomerDiscountV2 struct {
	// 折扣ID，唯一表示一条折扣信息。

	DiscountId *string `json:"discount_id,omitempty"`
	// 折扣率，精确到2位小数。 如果折扣率是85%，则折扣率写成0.85。

	Discount *float64 `json:"discount,omitempty"`
	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。

	EffectiveTime *string `json:"effective_time,omitempty"`
	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。

	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o QuerySubCustomerDiscountV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySubCustomerDiscountV2 struct{}"
	}

	return strings.Join([]string{"QuerySubCustomerDiscountV2", string(data)}, " ")
}
