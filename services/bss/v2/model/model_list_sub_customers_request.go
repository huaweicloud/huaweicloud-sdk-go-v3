package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomersRequest Request Object
type ListSubCustomersRequest struct {
	Body *QuerySubCustomerListReq `json:"body,omitempty"`
}

func (o ListSubCustomersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomersRequest struct{}"
	}

	return strings.Join([]string{"ListSubCustomersRequest", string(data)}, " ")
}
