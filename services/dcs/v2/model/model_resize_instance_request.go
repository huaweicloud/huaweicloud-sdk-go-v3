package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeInstanceRequest Request Object
type ResizeInstanceRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResizeInstanceBody `json:"body,omitempty"`
}

func (o ResizeInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceRequest struct{}"
	}

	return strings.Join([]string{"ResizeInstanceRequest", string(data)}, " ")
}
