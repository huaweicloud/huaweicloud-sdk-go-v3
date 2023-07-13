package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStructTemplateclusterInfo 结构化类型。
type ShowStructTemplateclusterInfo struct {

	// 测试
	ClusterName *string `json:"cluster_name,omitempty"`

	// 测试
	KafkaBootstrapServers *string `json:"kafka_bootstrap_servers,omitempty"`

	// 测试
	KafkaSslEnable *bool `json:"kafka_ssl_enable,omitempty"`
}

func (o ShowStructTemplateclusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStructTemplateclusterInfo struct{}"
	}

	return strings.Join([]string{"ShowStructTemplateclusterInfo", string(data)}, " ")
}
