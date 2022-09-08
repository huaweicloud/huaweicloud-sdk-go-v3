package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCustomersBalancesReq struct {

	// 客户信息列表。 具体请参见表1。
	CustomerInfos []CustomerInfoV2 `json:"customer_infos"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。  华为云总经销商（一级经销商）查询云经销商（二级经销商）子客户的账户余额时，需要携带该参数。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QueryCustomersBalancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCustomersBalancesReq struct{}"
	}

	return strings.Join([]string{"QueryCustomersBalancesReq", string(data)}, " ")
}
