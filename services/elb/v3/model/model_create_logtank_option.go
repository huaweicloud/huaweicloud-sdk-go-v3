package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogtankOption 创建云日志请求参数。
type CreateLogtankOption struct {

	// 负载均衡器id
	LoadbalancerId string `json:"loadbalancer_id"`

	// 日志组别id，其他（非ELB）服务提供
	LogGroupId string `json:"log_group_id"`

	// 日志订阅主题id，其他（非ELB）服务提供
	LogTopicId string `json:"log_topic_id"`
}

func (o CreateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankOption struct{}"
	}

	return strings.Join([]string{"CreateLogtankOption", string(data)}, " ")
}
