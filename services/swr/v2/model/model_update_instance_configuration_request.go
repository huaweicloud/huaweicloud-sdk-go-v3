package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationRequest Request Object
type UpdateInstanceConfigurationRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceConfigurationRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationRequest", string(data)}, " ")
}
