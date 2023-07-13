package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateColdVolumeRequest Request Object
type CreateColdVolumeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CreateColdVolumeRequestBody `json:"body,omitempty"`
}

func (o CreateColdVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateColdVolumeRequest struct{}"
	}

	return strings.Join([]string{"CreateColdVolumeRequest", string(data)}, " ")
}
