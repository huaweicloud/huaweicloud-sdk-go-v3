package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomersResponse Response Object
type ListSubCustomersResponse struct {

	// 子用户列表
	SubCustomerInfos *[]SubCutomerInfoV2 `json:"sub_customer_infos,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ListSubCustomersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomersResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomersResponse", string(data)}, " ")
}
