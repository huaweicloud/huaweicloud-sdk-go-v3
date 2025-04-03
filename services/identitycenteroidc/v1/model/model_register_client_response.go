package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClientResponse Response Object
type RegisterClientResponse struct {
	ClientInfo     *ClientInfoDto `json:"client_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o RegisterClientResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClientResponse struct{}"
	}

	return strings.Join([]string{"RegisterClientResponse", string(data)}, " ")
}
