package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 结构化类型。
type ShowStructTemplateclusterInfo struct {

	// 测试
	ClusterName *string `json:"cluster_name,omitempty" xml:"cluster_name"`

	// 测试
	KafkaBootstrapServers *string `json:"kafka_bootstrap_servers,omitempty" xml:"kafka_bootstrap_servers"`

	// 测试
	KafkaSslEnable *bool `json:"kafka_ssl_enable,omitempty" xml:"kafka_ssl_enable"`
}

func (o ShowStructTemplateclusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStructTemplateclusterInfo struct{}"
	}

	return strings.Join([]string{"ShowStructTemplateclusterInfo", string(data)}, " ")
}
