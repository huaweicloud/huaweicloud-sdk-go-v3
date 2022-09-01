package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyNode struct {

	// Proxy节点ID。
	Id string `json:"id" xml:"id"`

	// Proxy节点名称。
	Name string `json:"name" xml:"name"`

	// Proxy节点角色。 - master：主节点。 - slave：备节点。
	Role string `json:"role" xml:"role"`

	// 可用区。
	AzCode string `json:"az_code" xml:"az_code"`

	// Proxy节点状态。 - normal：正常。 - abnormal：异常。 - creating：创建中。 - deleted：已删除。
	Status string `json:"status" xml:"status"`

	// Proxy节点是否被冻结。 - 0：未冻结。 - 1：冻结。 - 2：冻结删除。
	FrozenFlag int32 `json:"frozen_flag" xml:"frozen_flag"`
}

func (o ProxyNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyNode struct{}"
	}

	return strings.Join([]string{"ProxyNode", string(data)}, " ")
}
