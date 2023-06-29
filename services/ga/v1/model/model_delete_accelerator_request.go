package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAcceleratorRequest Request Object
type DeleteAcceleratorRequest struct {

	// 全球加速器ID。
	AcceleratorId string `json:"accelerator_id"`
}

func (o DeleteAcceleratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAcceleratorRequest struct{}"
	}

	return strings.Join([]string{"DeleteAcceleratorRequest", string(data)}, " ")
}
