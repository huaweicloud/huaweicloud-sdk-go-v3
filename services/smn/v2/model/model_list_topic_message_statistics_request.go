package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopicMessageStatisticsRequest Request Object
type ListTopicMessageStatisticsRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// 起始时间
	StartTime string `json:"start_time"`

	// 结束时间
	EndTime string `json:"end_time"`
}

func (o ListTopicMessageStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicMessageStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListTopicMessageStatisticsRequest", string(data)}, " ")
}
