package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExchangeInstanceIpResponse Response Object
type ExchangeInstanceIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExchangeInstanceIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExchangeInstanceIpResponse struct{}"
	}

	return strings.Join([]string{"ExchangeInstanceIpResponse", string(data)}, " ")
}
