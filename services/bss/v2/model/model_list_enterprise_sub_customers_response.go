package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEnterpriseSubCustomersResponse struct {

	// 结果集数量，成功才有。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 客户信息列表，成功才有。 具体请参见表2。
	SubCustomerInfos *[]SubCustomerInfoV2 `json:"sub_customer_infos,omitempty" xml:"sub_customer_infos"`
	HttpStatusCode   int                  `json:"-"`
}

func (o ListEnterpriseSubCustomersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseSubCustomersResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseSubCustomersResponse", string(data)}, " ")
}
