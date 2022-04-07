package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDocWatermarkByAddressRequest struct {
	Body *ShowDocWatermarkByAddressRequestBody `json:"body,omitempty"`
}

func (o ShowDocWatermarkByAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDocWatermarkByAddressRequest struct{}"
	}

	return strings.Join([]string{"ShowDocWatermarkByAddressRequest", string(data)}, " ")
}
