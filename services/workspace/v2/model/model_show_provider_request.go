package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProviderRequest Request Object
type ShowProviderRequest struct {

	// 供应商id。
	ProviderId string `json:"provider_id"`
}

func (o ShowProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProviderRequest struct{}"
	}

	return strings.Join([]string{"ShowProviderRequest", string(data)}, " ")
}
