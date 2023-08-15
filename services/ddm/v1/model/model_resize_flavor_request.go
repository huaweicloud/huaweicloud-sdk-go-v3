package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeFlavorRequest Request Object
type ResizeFlavorRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResizeFlavorReq `json:"body,omitempty"`
}

func (o ResizeFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeFlavorRequest struct{}"
	}

	return strings.Join([]string{"ResizeFlavorRequest", string(data)}, " ")
}
