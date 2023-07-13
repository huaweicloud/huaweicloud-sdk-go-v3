package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAcceleratorResponse Response Object
type CreateAcceleratorResponse struct {
	Accelerator *AcceleratorDetail `json:"accelerator,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAcceleratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAcceleratorResponse struct{}"
	}

	return strings.Join([]string{"CreateAcceleratorResponse", string(data)}, " ")
}
