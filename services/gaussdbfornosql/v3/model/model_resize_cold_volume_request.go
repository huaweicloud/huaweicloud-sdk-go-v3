package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeColdVolumeRequest Request Object
type ResizeColdVolumeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResizeColdVolumeRequestBody `json:"body,omitempty"`
}

func (o ResizeColdVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeColdVolumeRequest struct{}"
	}

	return strings.Join([]string{"ResizeColdVolumeRequest", string(data)}, " ")
}
