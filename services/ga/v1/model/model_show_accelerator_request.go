package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAcceleratorRequest Request Object
type ShowAcceleratorRequest struct {

	// 全球加速器ID。
	AcceleratorId string `json:"accelerator_id"`
}

func (o ShowAcceleratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAcceleratorRequest struct{}"
	}

	return strings.Join([]string{"ShowAcceleratorRequest", string(data)}, " ")
}
