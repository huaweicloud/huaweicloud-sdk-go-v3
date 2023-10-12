package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationDataMetadata 转发策略类型访问方式的负载均衡策略配置。
type ConfigurationDataMetadata struct {

	// 如下配置表示：负载均衡分配策略使用加权轮询算法，不启用健康检查。 - \"kubernetes.io/elb.health-check-flag\": \"off\" - \"kubernetes.io/elb.lb-algorithm\": \"ROUND_ROBIN\"
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o ConfigurationDataMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDataMetadata struct{}"
	}

	return strings.Join([]string{"ConfigurationDataMetadata", string(data)}, " ")
}
