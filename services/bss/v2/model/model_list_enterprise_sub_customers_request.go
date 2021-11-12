package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEnterpriseSubCustomersRequest struct {
	// 企业子账号的账号名。根据fuzzy_query取值决定是否按模糊查询。

	SubCustomerAccountName *string `json:"sub_customer_account_name,omitempty"`
	// 企业子账号的显示名称。不限制特殊字符。根据fuzzy_query取值决定是否按模糊查询。

	SubCustomerDisplayName *string `json:"sub_customer_display_name,omitempty"`
	// 企业子账号的显示名称、用户名是否按模糊查询。0：不按模糊查询1：按模糊查询默认值为0。

	FuzzyQuery *int32 `json:"fuzzy_query,omitempty"`
	// 偏移量，从0开始，默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询记录数，默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 子账号归属的组织单元ID。

	OrgId *string `json:"org_id,omitempty"`
}

func (o ListEnterpriseSubCustomersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseSubCustomersRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseSubCustomersRequest", string(data)}, " ")
}
