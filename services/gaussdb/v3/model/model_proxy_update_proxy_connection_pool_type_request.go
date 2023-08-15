package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProxyUpdateProxyConnectionPoolTypeRequest proxy更新连接池请求体。
type ProxyUpdateProxyConnectionPoolTypeRequest struct {

	// 连接池类型。 - CLOSED：关闭连接池。 - SESSION：开启会话级连接池
	ConnectionPoolType string `json:"connection_pool_type"`
}

func (o ProxyUpdateProxyConnectionPoolTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyUpdateProxyConnectionPoolTypeRequest struct{}"
	}

	return strings.Join([]string{"ProxyUpdateProxyConnectionPoolTypeRequest", string(data)}, " ")
}
