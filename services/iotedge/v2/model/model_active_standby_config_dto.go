package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 高可用主备配置请求结构体
type ActiveStandbyConfigDto struct {

	// 当前的工作节点，主节点还是备节点在工作，初始创建时工作节点为DEFAULT节点，(DEFAULT|MASTER|SLAVE)
	WorkNode *string `json:"work_node,omitempty"`

	// 主节点网卡名称
	MasterInterfaceName *string `json:"master_interface_name,omitempty"`

	// 备节点网卡名称
	SlaveInterfaceName *string `json:"slave_interface_name,omitempty"`

	// 网卡ip
	VirtualIpAddress *string `json:"virtual_ip_address,omitempty"`
}

func (o ActiveStandbyConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActiveStandbyConfigDto struct{}"
	}

	return strings.Join([]string{"ActiveStandbyConfigDto", string(data)}, " ")
}
