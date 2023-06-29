package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTerminalsBindingDesktopsConfigResponse Response Object
type UpdateTerminalsBindingDesktopsConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTerminalsBindingDesktopsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTerminalsBindingDesktopsConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateTerminalsBindingDesktopsConfigResponse", string(data)}, " ")
}
