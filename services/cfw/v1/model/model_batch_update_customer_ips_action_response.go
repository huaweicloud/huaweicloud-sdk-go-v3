package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateCustomerIpsActionResponse Response Object
type BatchUpdateCustomerIpsActionResponse struct {
	Data           *BatchCustomerIpsRespData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o BatchUpdateCustomerIpsActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateCustomerIpsActionResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateCustomerIpsActionResponse", string(data)}, " ")
}
