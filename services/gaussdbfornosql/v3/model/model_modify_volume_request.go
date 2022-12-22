package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ModifyVolumeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyVolumeRequestBody `json:"body,omitempty"`
}

func (o ModifyVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVolumeRequest struct{}"
	}

	return strings.Join([]string{"ModifyVolumeRequest", string(data)}, " ")
}
