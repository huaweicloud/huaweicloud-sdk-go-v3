package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigurationAggregatorRequest Request Object
type CreateConfigurationAggregatorRequest struct {
	Body *ConfigurationAggregatorRequest `json:"body,omitempty"`
}

func (o CreateConfigurationAggregatorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationAggregatorRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationAggregatorRequest", string(data)}, " ")
}
