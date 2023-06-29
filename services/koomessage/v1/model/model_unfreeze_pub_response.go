package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnfreezePubResponse Response Object
type UnfreezePubResponse struct {
	Data           *UnfreezePubResponseModel `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o UnfreezePubResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezePubResponse struct{}"
	}

	return strings.Join([]string{"UnfreezePubResponse", string(data)}, " ")
}
