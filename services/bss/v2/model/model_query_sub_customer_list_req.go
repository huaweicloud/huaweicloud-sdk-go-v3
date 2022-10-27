package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuerySubCustomerListReq struct {

	// 客户登录名称（如果客户创建了IAM用户，此处需要填写主账号登录名称。关于主账号和IAM用户的具体介绍请参见身份管理中“账号”和“IAM用户”的描述）。 支持模糊查询。仅支持前缀匹配、后缀匹配、中间匹配；不支持携带空格查询。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	AccountName *string `json:"account_name,omitempty"`

	// 客户的实名认证名称，支持模糊查询。仅支持前缀匹配、后缀匹配、中间匹配；不支持携带空格查询。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	Customer *string `json:"customer,omitempty"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的客户数量。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 标签，支持模糊查找。仅支持前缀匹配、后缀匹配、中间匹配；不支持携带空格查询；不支持英文大小写模糊匹配查询。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	Label *string `json:"label,omitempty"`

	// 关联类型： 1：顾问销售2：代售 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	AssociationType *string `json:"association_type,omitempty"`

	// 关联时间区间段开始，UTC时间。 格式：YYYY-MM-DD'T'hh:mm:ss'Z'，例如“2019-05-06T08:05:01Z”。 此参数不携带或携带值为null时，不作为筛选条件，不支持携带值为空串。
	AssociatedOnBegin *string `json:"associated_on_begin,omitempty"`

	// 关联时间区间段结束，UTC时间。 格式：YYYY-MM-DD'T'hh:mm:ss'Z'，例如“2019-05-06T08:05:01Z”。 此参数不携带或携带值为null时，不作为筛选条件，不支持携带值为空串。
	AssociatedOnEnd *string `json:"associated_on_end,omitempty"`

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id，或者可以从创建客户接口的响应获取domain_id。此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	CustomerId *string `json:"customer_id,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。如果需要查询云经销商的子客户列表，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QuerySubCustomerListReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySubCustomerListReq struct{}"
	}

	return strings.Join([]string{"QuerySubCustomerListReq", string(data)}, " ")
}
