package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMultiAccountRetrieveCouponsRequest Request Object
type ListMultiAccountRetrieveCouponsRequest struct {

	// 企业子账户的账号ID。
	SubCustomerId string `json:"sub_customer_id"`

	// 偏移量，默认值为0。此参数不携带或携带值为空或携带值为null时，取默认值0；不支持携带值为空串。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询条数，默认值为10。此参数不携带或携带值为空或携带值为null时，取默认值0；不支持携带值为空串。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMultiAccountRetrieveCouponsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMultiAccountRetrieveCouponsRequest struct{}"
	}

	return strings.Join([]string{"ListMultiAccountRetrieveCouponsRequest", string(data)}, " ")
}
