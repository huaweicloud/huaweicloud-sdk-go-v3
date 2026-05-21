package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserDefinedDomainConfigRequestBody **参数解释**：自定义负载均衡器域名解析的请求参数。
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
