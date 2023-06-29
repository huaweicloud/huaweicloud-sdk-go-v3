package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLatelyApiStatisticsV2Request Request Object
type ListLatelyApiStatisticsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// API的编号
	ApiId string `json:"api_id"`

	// 最近统计时长，单位必须为h和m，比如1h和1m，分别代表最近1小时和最近1分钟
	Duration string `json:"duration"`
}

func (o ListLatelyApiStatisticsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLatelyApiStatisticsV2Request struct{}"
	}

	return strings.Join([]string{"ListLatelyApiStatisticsV2Request", string(data)}, " ")
}
