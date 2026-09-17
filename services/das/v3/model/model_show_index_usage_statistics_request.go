package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIndexUsageStatisticsRequest Request Object
type ShowIndexUsageStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowIndexUsageStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndexUsageStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowIndexUsageStatisticsRequest", string(data)}, " ")
}
