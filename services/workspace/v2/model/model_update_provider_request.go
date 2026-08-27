package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProviderRequest Request Object
type UpdateProviderRequest struct {

	// 供应商id。
	ProviderId string `json:"provider_id"`

	Body *UpdateProviderReq `json:"body,omitempty"`
}

func (o UpdateProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProviderRequest struct{}"
	}

	return strings.Join([]string{"UpdateProviderRequest", string(data)}, " ")
}
