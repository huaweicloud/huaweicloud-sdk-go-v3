package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateElasticResourcePoolQueueScalingPolicyInfo 更新弹性资源池队列消息请求
type UpdateElasticResourcePoolQueueScalingPolicyInfo struct {

	// 该队列在该弹性资源池下的扩缩容策略信息。单条策略信息包含时间段、优先级和CU范围。 每个队列至少要配置一条时间段为[00:00, 24:00]的默认扩缩容策略。
	QueueScalingPolicies []QueueScalingPolicyInfo `json:"queue_scaling_policies"`
}

func (o UpdateElasticResourcePoolQueueScalingPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateElasticResourcePoolQueueScalingPolicyInfo struct{}"
	}

	return strings.Join([]string{"UpdateElasticResourcePoolQueueScalingPolicyInfo", string(data)}, " ")
}
