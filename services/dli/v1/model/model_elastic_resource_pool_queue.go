package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ElasticResourcePoolQueue struct {

	// 队列名称
	QueueName *string `json:"queue_name,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 队列类型
	QueueType *string `json:"queue_type,omitempty"`

	// 队列扩缩容策略
	QueueScalingPolicies *[]QueueScalingPolicy `json:"queue_scaling_policies,omitempty"`

	// 所有者
	Owner *string `json:"owner,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 引擎类型
	Engine *string `json:"engine,omitempty"`
}

func (o ElasticResourcePoolQueue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElasticResourcePoolQueue struct{}"
	}

	return strings.Join([]string{"ElasticResourcePoolQueue", string(data)}, " ")
}
