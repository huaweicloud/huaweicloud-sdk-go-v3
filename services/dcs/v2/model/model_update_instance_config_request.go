package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigRequest Request Object
type UpdateInstanceConfigRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ModifyRedisConfigBody `json:"body,omitempty"`
}

func (o UpdateInstanceConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigRequest", string(data)}, " ")
}
