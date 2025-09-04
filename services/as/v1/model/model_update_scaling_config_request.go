package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScalingConfigRequest Request Object
type UpdateScalingConfigRequest struct {
	ScalingConfigurationId string `json:"scaling_configuration_id"`

	Body *UpdateScalingConfigOption `json:"body,omitempty"`
}

func (o UpdateScalingConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateScalingConfigRequest", string(data)}, " ")
}
