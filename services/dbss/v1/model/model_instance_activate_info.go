package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceActivateInfo 实例激活信息
type InstanceActivateInfo struct {

	// 代理DOMAIN名称
	DscProxyDomainName *string `json:"dsc_proxy_domain_name,omitempty"`

	// 代理IP
	DscProxyIp *string `json:"dsc_proxy_ip,omitempty"`

	// 代理端口
	DscProxyPort *int32 `json:"dsc_proxy_port,omitempty"`

	// 主节点ID
	MasterNodeId *string `json:"master_node_id,omitempty"`

	// 备节点ID
	SlaveNodeId *string `json:"slave_node_id,omitempty"`
}

func (o InstanceActivateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceActivateInfo struct{}"
	}

	return strings.Join([]string{"InstanceActivateInfo", string(data)}, " ")
}
