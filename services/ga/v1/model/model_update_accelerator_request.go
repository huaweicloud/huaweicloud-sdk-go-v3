package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAcceleratorRequest Request Object
type UpdateAcceleratorRequest struct {

	// 全球加速器ID。
	AcceleratorId string `json:"accelerator_id"`

	Body *UpdateAcceleratorRequestBody `json:"body,omitempty"`
}

func (o UpdateAcceleratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAcceleratorRequest struct{}"
	}

	return strings.Join([]string{"UpdateAcceleratorRequest", string(data)}, " ")
}
