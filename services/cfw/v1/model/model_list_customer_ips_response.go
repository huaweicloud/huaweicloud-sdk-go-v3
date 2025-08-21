package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerIpsResponse Response Object
type ListCustomerIpsResponse struct {
	Data           *CustomerIpsPageInfo `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListCustomerIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerIpsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerIpsResponse", string(data)}, " ")
}
