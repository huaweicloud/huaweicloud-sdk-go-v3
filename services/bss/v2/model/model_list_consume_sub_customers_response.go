package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConsumeSubCustomersResponse Response Object
type ListConsumeSubCustomersResponse struct {

	// 子客户列表。 具体请参见表SubCustomerInfo。
	SubCustomers *[]SubCustomerInfoV3 `json:"sub_customers,omitempty"`

	// 结果集数量，只有成功才返回这个参数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListConsumeSubCustomersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumeSubCustomersResponse struct{}"
	}

	return strings.Join([]string{"ListConsumeSubCustomersResponse", string(data)}, " ")
}
