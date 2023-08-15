package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDevicesRequest Request Object
type BatchDeleteDevicesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *BatchDeleteDevicesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDevicesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDevicesRequest", string(data)}, " ")
}
