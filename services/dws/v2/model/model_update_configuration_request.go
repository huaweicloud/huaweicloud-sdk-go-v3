package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConfigurationRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`

	// 参数组ID。
	ConfigurationId string `json:"configuration_id"`

	Body *ConfigurationParameterValues `json:"body,omitempty"`
}

func (o UpdateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationRequest", string(data)}, " ")
}
