package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowImageWatermarkByAddressRequest struct {
	Body *ShowImageWatermarkByAddressRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ShowImageWatermarkByAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkByAddressRequest struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkByAddressRequest", string(data)}, " ")
}
