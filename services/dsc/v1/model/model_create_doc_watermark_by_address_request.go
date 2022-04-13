package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDocWatermarkByAddressRequest struct {
	Body *CreateDocWatermarkByAddressRequestBody `json:"body,omitempty"`
}

func (o CreateDocWatermarkByAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocWatermarkByAddressRequest struct{}"
	}

	return strings.Join([]string{"CreateDocWatermarkByAddressRequest", string(data)}, " ")
}
