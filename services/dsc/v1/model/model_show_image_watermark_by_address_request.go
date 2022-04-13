package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowImageWatermarkByAddressRequest struct {
	Body *ShowImageWatermarkByAddressRequestBody `json:"body,omitempty"`
}

func (o ShowImageWatermarkByAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkByAddressRequest struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkByAddressRequest", string(data)}, " ")
}
