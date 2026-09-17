package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIndexUsageTrendRequest Request Object
type ShowIndexUsageTrendRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowIndexUsageTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndexUsageTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowIndexUsageTrendRequest", string(data)}, " ")
}
