package model

import (
	"encoding/json"

	"strings"
)

type QuerySubCustomerListReq struct {
	// 客户登录名称（如果客户创建了IAM用户，此处需要填写主账号登录名称。关于主账号和IAM用户的具体介绍请参见身份管理中“账号”和“IAM用户”的描述）。 支持模糊查询。

	AccountName *string `json:"account_name,omitempty"`
	// 客户的实名认证名称，支持模糊查询。

	Customer *string `json:"customer,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的客户数量。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 标签，支持模糊查找。

	Label *string `json:"label,omitempty"`
	// 关联类型： 1：推荐2：垫付

	AssociationType *string `json:"association_type,omitempty"`
	// 关联时间区间段开始，UTC时间。 格式：YYYY-MM-DD'T'hh:mm:ss'Z'，例如“2019-05-06T08:05:01Z”。

	AssociatedOnBegin *string `json:"associated_on_begin,omitempty"`
	// 关联时间区间段结束，UTC时间。 格式：YYYY-MM-DD'T'hh:mm:ss'Z'，例如“2019-05-06T08:05:01Z”。

	AssociatedOnEnd *string `json:"associated_on_end,omitempty"`
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId *string `json:"customer_id,omitempty"`
	// 精英服务商ID。如果需要查询精英服务商伙伴的子客户列表，必须携带该字段。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QuerySubCustomerListReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuerySubCustomerListReq struct{}"
	}

	return strings.Join([]string{"QuerySubCustomerListReq", string(data)}, " ")
}
