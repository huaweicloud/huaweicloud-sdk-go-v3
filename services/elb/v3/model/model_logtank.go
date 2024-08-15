package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Logtank 云日志信息。
type Logtank struct {

	// 云日志ID。
	Id string `json:"id"`

	// 参数解释：项目ID。
	ProjectId string `json:"project_id"`

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	// 云日志分组ID。
	LogGroupId string `json:"log_group_id"`

	// 云日志主题ID。
	LogTopicId string `json:"log_topic_id"`
}

func (o Logtank) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Logtank struct{}"
	}

	return strings.Join([]string{"Logtank", string(data)}, " ")
}
