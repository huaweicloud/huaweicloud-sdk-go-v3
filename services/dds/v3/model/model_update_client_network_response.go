package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateClientNetworkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateClientNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientNetworkResponse struct{}"
	}

	return strings.Join([]string{"UpdateClientNetworkResponse", string(data)}, " ")
}
