package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResizeInstanceFlavorRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *OpenGaussResizeRequest `json:"body,omitempty"`
}

func (o ResizeInstanceFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceFlavorRequest struct{}"
	}

	return strings.Join([]string{"ResizeInstanceFlavorRequest", string(data)}, " ")
}
