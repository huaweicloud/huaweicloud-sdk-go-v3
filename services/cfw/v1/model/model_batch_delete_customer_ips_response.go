package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCustomerIpsResponse Response Object
type BatchDeleteCustomerIpsResponse struct {
	Data           *BatchCustomerIpsRespData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o BatchDeleteCustomerIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCustomerIpsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCustomerIpsResponse", string(data)}, " ")
}
