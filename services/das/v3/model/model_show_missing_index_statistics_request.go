package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMissingIndexStatisticsRequest Request Object
type ShowMissingIndexStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowMissingIndexStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMissingIndexStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowMissingIndexStatisticsRequest", string(data)}, " ")
}
