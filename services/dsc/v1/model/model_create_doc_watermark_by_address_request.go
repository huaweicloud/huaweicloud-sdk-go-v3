package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDocWatermarkByAddressRequest Request Object
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
