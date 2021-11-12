package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLatelyGroupStatisticsV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API分组的编号

	GroupId string `json:"group_id"`
}

func (o ListLatelyGroupStatisticsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLatelyGroupStatisticsV2Request struct{}"
	}

	return strings.Join([]string{"ListLatelyGroupStatisticsV2Request", string(data)}, " ")
}
