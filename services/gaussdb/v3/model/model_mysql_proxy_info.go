package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Proxy信息。
type MysqlProxyInfo struct {

	// Proxy实例ID。
	PoolId *string `json:"pool_id,omitempty"`

	// Proxy实例名称。
	Name *string `json:"name,omitempty"`

	// Proxy读写分离地址。
	Address *string `json:"address,omitempty"`
}

func (o MysqlProxyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxyInfo struct{}"
	}

	return strings.Join([]string{"MysqlProxyInfo", string(data)}, " ")
}
