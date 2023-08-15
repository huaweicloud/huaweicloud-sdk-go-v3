package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationAggregatorRequest Request Object
type ShowConfigurationAggregatorRequest struct {

	// 资源聚合器ID。
	AggregatorId string `json:"aggregator_id"`
}

func (o ShowConfigurationAggregatorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationAggregatorRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationAggregatorRequest", string(data)}, " ")
}
