package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchFreezeDevicesRequest Request Object
type BatchFreezeDevicesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *BatchFreezeDevicesRequestBody `json:"body,omitempty"`
}

func (o BatchFreezeDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchFreezeDevicesRequest struct{}"
	}

	return strings.Join([]string{"BatchFreezeDevicesRequest", string(data)}, " ")
}
