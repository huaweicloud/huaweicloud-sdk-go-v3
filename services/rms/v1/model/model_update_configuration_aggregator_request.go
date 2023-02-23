package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConfigurationAggregatorRequest struct {

	// 资源聚合器ID。
	AggregatorId string `json:"aggregator_id"`

	Body *ConfigurationAggregatorRequest `json:"body,omitempty"`
}

func (o UpdateConfigurationAggregatorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationAggregatorRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationAggregatorRequest", string(data)}, " ")
}
