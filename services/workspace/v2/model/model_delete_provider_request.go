package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProviderRequest Request Object
type DeleteProviderRequest struct {

	// 供应商id。
	ProviderId string `json:"provider_id"`
}

func (o DeleteProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProviderRequest struct{}"
	}

	return strings.Join([]string{"DeleteProviderRequest", string(data)}, " ")
}
