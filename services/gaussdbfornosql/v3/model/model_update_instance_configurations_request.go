package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationsRequest Request Object
type UpdateInstanceConfigurationsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceConfigurationsRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationsRequest", string(data)}, " ")
}
