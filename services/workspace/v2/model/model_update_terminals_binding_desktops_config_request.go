package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTerminalsBindingDesktopsConfigRequest Request Object
type UpdateTerminalsBindingDesktopsConfigRequest struct {
	Body *TerminalsBindingDesktopsConfig `json:"body,omitempty"`
}

func (o UpdateTerminalsBindingDesktopsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTerminalsBindingDesktopsConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateTerminalsBindingDesktopsConfigRequest", string(data)}, " ")
}
