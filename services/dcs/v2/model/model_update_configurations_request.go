package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConfigurationsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ModifyRedisConfigBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationsRequest", string(data)}, " ")
}
