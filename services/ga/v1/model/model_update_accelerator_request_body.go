package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// update Accelerator request
type UpdateAcceleratorRequestBody struct {
	Accelerator *UpdateAcceleratorOption `json:"accelerator"`
}

func (o UpdateAcceleratorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAcceleratorRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAcceleratorRequestBody", string(data)}, " ")
}
