package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAcceleratorResponse Response Object
type ShowAcceleratorResponse struct {
	Accelerator *AcceleratorDetail `json:"accelerator,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAcceleratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAcceleratorResponse struct{}"
	}

	return strings.Join([]string{"ShowAcceleratorResponse", string(data)}, " ")
}
