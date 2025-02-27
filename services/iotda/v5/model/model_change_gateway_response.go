package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeGatewayResponse Response Object
type ChangeGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeGatewayResponse struct{}"
	}

	return strings.Join([]string{"ChangeGatewayResponse", string(data)}, " ")
}
