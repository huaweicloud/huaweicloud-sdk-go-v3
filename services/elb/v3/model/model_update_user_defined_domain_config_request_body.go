package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserDefinedDomainConfigRequestBody **参数解释**：配置负载均衡器用户自定义域名化相关配置参数。
type UpdateUserDefinedDomainConfigRequestBody struct {
	Loadbalancer *UserDefinedDnsConfigRequestBody `json:"loadbalancer"`
}

func (o UpdateUserDefinedDomainConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserDefinedDomainConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateUserDefinedDomainConfigRequestBody", string(data)}, " ")
}
