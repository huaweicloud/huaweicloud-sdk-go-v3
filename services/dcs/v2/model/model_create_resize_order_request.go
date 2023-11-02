package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResizeOrderRequest Request Object
type CreateResizeOrderRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateResizeOrderRequestBody `json:"body,omitempty"`
}

func (o CreateResizeOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResizeOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateResizeOrderRequest", string(data)}, " ")
}
