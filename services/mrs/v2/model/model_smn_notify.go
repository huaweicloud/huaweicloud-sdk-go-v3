package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmnNotify smn消息订阅信息。用于订阅集群全量告警信息。如果需要配置更精细化的告警或者事件，请于创建集群完成后，在”告警管理“中进行设置。
type SmnNotify struct {

	// SMN消息通知服务的主题urn，如果需要开启告警订阅，则必填。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 该订阅规则名称。如果不填写，则默认为default_alert_rule。
	SubscriptionName *string `json:"subscription_name,omitempty"`
}

func (o SmnNotify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnNotify struct{}"
	}

	return strings.Join([]string{"SmnNotify", string(data)}, " ")
}
