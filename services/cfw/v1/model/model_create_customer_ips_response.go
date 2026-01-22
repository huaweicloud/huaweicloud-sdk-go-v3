package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomerIpsResponse Response Object
type CreateCustomerIpsResponse struct {
	Data           *CreateCustomerIpsRespData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CreateCustomerIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomerIpsResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomerIpsResponse", string(data)}, " ")
}
