package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConfigurationMetadata 访问方式配置元数据。
type AccessConfigurationMetadata struct {

	// 附加参数。  举例：负载均衡分配策略使用加权轮询算法，不启用健康检查。配置如下： - \"kubernetes.io/elb.health-check-flag\": \"off\" - \"kubernetes.io/elb.lb-algorithm\": \"ROUND_ROBIN\"
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o AccessConfigurationMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigurationMetadata struct{}"
	}

	return strings.Join([]string{"AccessConfigurationMetadata", string(data)}, " ")
}
