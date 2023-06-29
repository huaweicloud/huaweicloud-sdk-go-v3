package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomersResponse Response Object
type ListSubCustomersResponse struct {

	// 客户信息列表。 具体请参见表1。
	CustomerInfos *[]CustomerInformation `json:"customer_infos,omitempty"`

	// 总记录数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubCustomersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomersResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomersResponse", string(data)}, " ")
}
