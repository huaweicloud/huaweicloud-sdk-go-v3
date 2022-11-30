package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLatelyGroupStatisticsV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
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
