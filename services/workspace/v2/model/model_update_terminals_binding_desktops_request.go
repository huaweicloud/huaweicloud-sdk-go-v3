package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTerminalsBindingDesktopsRequest struct {
	Body *UpdateTerminalsBindingDesktopsRequestBody `json:"body,omitempty"`
}

func (o UpdateTerminalsBindingDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTerminalsBindingDesktopsRequest struct{}"
	}

	return strings.Join([]string{"UpdateTerminalsBindingDesktopsRequest", string(data)}, " ")
}
