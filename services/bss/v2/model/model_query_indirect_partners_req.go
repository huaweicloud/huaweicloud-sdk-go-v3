package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryIndirectPartnersReq struct {

	// 精英服务商伙伴的账号名。
	AccountName *string `json:"account_name,omitempty"`

	// 精英服务商关联华为云伙伴能力中心的开始时间。 UTC时间（包括时区），比如2016-03-28T00:00:00Z
	AssociatedOnBegin *string `json:"associated_on_begin,omitempty"`

	// 精英服务商关联华为云伙伴能力中心的结束时间。 UTC时间（包括时区），比如2016-03-28T00:00:00Z
	AssociatedOnEnd *string `json:"associated_on_end,omitempty"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量限制。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 精英服务商ID。获取方法请参见[查询精英服务商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QueryIndirectPartnersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryIndirectPartnersReq struct{}"
	}

	return strings.Join([]string{"QueryIndirectPartnersReq", string(data)}, " ")
}
