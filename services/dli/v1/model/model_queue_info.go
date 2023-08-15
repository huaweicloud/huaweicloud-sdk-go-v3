package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueueInfo struct {

	// 队列名称
	QueueName *string `json:"queue_name,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 队列类型
	QueueType *string `json:"queue_type,omitempty"`

	// 队列扩缩容策略
	QueueScalingPolicies *[]QueueScalingPoliciesResponse `json:"queue_scaling_policies,omitempty"`

	// 所有者
	Owner *string `json:"owner,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 引擎类型
	Engine *string `json:"engine,omitempty"`
}

func (o QueueInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueInfo struct{}"
	}

	return strings.Join([]string{"QueueInfo", string(data)}, " ")
}
