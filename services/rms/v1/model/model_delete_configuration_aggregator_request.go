package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteConfigurationAggregatorRequest struct {

	// 资源聚合器ID。
	AggregatorId string `json:"aggregator_id"`
}

func (o DeleteConfigurationAggregatorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigurationAggregatorRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigurationAggregatorRequest", string(data)}, " ")
}
