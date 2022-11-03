package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAcceleratorResponse struct {
	Accelerator *AcceleratorDetail `json:"accelerator,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAcceleratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAcceleratorResponse struct{}"
	}

	return strings.Join([]string{"UpdateAcceleratorResponse", string(data)}, " ")
}
