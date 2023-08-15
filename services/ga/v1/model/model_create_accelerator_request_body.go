package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAcceleratorRequestBody create Accelerator request
type CreateAcceleratorRequestBody struct {
	Accelerator *CreateAcceleratorOption `json:"accelerator"`
}

func (o CreateAcceleratorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAcceleratorRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAcceleratorRequestBody", string(data)}, " ")
}
