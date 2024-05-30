package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProxyUpdateProxyNameRequest 修改代理实例名称的请求体对象。
type ProxyUpdateProxyNameRequest struct {

	// 修改代理实例的新名称
	Alias string `json:"alias"`
}

func (o ProxyUpdateProxyNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyUpdateProxyNameRequest struct{}"
	}

	return strings.Join([]string{"ProxyUpdateProxyNameRequest", string(data)}, " ")
}
