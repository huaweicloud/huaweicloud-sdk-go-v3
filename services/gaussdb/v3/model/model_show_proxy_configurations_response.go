package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyConfigurationsResponse Response Object
type ShowProxyConfigurationsResponse struct {

	// 数据总数
	TotalCount *string `json:"total_count,omitempty"`

	// 内核可配置的参数列表
	Configurations *[]ProxyConfiguration `json:"configurations,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowProxyConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ShowProxyConfigurationsResponse", string(data)}, " ")
}
