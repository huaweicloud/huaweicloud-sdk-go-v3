package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerNewTagResponse Response Object
type ListSubCustomerNewTagResponse struct {

	// 新客标签信息列表。只有成功的时候才返回。 此列表不包含非子客户的数据。批量查询客户新客标签时，若入参携带了非子客户，会被过滤。 具体请参见表1 NewCustomerTagItem。
	NewCustomerTags *[]NewCustomerTagItem `json:"new_customer_tags,omitempty"`
	HttpStatusCode  int                   `json:"-"`
}

func (o ListSubCustomerNewTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerNewTagResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomerNewTagResponse", string(data)}, " ")
}
