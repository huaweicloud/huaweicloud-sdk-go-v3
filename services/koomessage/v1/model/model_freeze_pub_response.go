package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type FreezePubResponse struct {
	Data           *FreezePubResponseModel `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o FreezePubResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezePubResponse struct{}"
	}

	return strings.Join([]string{"FreezePubResponse", string(data)}, " ")
}
