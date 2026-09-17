package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportSubscription 健康报告订阅信息
type ReportSubscription struct {

	// 订阅ID
	SubscribeId *string `json:"subscribe_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 租户在某一Region下的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 协议
	Protocol *string `json:"protocol,omitempty"`

	// 地址
	Endpoint *string `json:"endpoint,omitempty"`

	// 主题
	Topic *string `json:"topic,omitempty"`

	// 主题地址
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 桶名
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// 风险等级
	Level *string `json:"level,omitempty"`

	// 语言
	Locale *string `json:"locale,omitempty"`

	// 额外信息
	Extra *interface{} `json:"extra,omitempty"`
}

func (o ReportSubscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportSubscription struct{}"
	}

	return strings.Join([]string{"ReportSubscription", string(data)}, " ")
}
