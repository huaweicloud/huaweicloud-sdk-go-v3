package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTerminalsBindingDesktopsResponse Response Object
type UpdateTerminalsBindingDesktopsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTerminalsBindingDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTerminalsBindingDesktopsResponse struct{}"
	}

	return strings.Join([]string{"UpdateTerminalsBindingDesktopsResponse", string(data)}, " ")
}
