package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackExchangeInstanceIpResponse Response Object
type RollbackExchangeInstanceIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RollbackExchangeInstanceIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackExchangeInstanceIpResponse struct{}"
	}

	return strings.Join([]string{"RollbackExchangeInstanceIpResponse", string(data)}, " ")
}
