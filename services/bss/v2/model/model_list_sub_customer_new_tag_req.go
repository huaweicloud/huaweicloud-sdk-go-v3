package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubCustomerNewTagReq struct {

	// 客户ID列表。单个客户ID最大长度：64。 此参数不携带或携带值为空列表或携带值为null时，不作为筛选条件。
	CustomerIds *[]string `json:"customer_ids,omitempty"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 示例1，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。 示例2，查询总数20条，期望每页返回10条数据，则获取第一页数据，入参offset填写0，limit填写10；获取第二页数据，入参offset填写10，limit填写10。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的客户数量。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 云经销商ID。获取方法请参见查询云经销商列表。如果需要查询云经销商的客户新客标签，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListSubCustomerNewTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerNewTagReq struct{}"
	}

	return strings.Join([]string{"ListSubCustomerNewTagReq", string(data)}, " ")
}
