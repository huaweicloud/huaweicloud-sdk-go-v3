package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageWatermarkByAddressRequest Request Object
type CreateImageWatermarkByAddressRequest struct {
	Body *CreateImageWatermarkByAddressRequestBody `json:"body,omitempty"`
}

func (o CreateImageWatermarkByAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkByAddressRequest struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkByAddressRequest", string(data)}, " ")
}
