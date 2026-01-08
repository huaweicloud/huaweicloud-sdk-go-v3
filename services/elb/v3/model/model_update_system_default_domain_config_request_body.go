package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSystemDefaultDomainConfigRequestBody **参数解释**：配置负载均衡器系统默认域名化相关配置参数。
type UpdateSystemDefaultDomainConfigRequestBody struct {
	Loadbalancer *SystemDefaultDnsConfigRequestBody `json:"loadbalancer"`
}

func (o UpdateSystemDefaultDomainConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSystemDefaultDomainConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSystemDefaultDomainConfigRequestBody", string(data)}, " ")
}
