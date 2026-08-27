package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnlargeProxyRequest proxy节点扩容信息
type EnlargeProxyRequest struct {

	// proxy节点扩容操作需要扩容的节点数。  扩容的节点数的取值范围：1~30之间的整数。  限制条件：该实例的proxy节点的总数量小于等于32。
	NodeNum int32 `json:"node_num"`

	// 数据库代理ID。  如果实例只开启了一个代理，可不传该参数；如果实例开启了多个代理，则必须指定一个数据库代理，扩容新的代理节点。
	ProxyId *string `json:"proxy_id,omitempty"`

	// **参数解释**：  数据库代理节点的可用区设置。  **约束限制**：  不传该字段，代理节点可用区将随机设置，优先与数据库节点可用区保持一致；传入该字段，代理节点将设置在指定可用区。
	ProxyNodesAzList *[]string `json:"proxy_nodes_az_list,omitempty"`
}

func (o EnlargeProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeProxyRequest struct{}"
	}

	return strings.Join([]string{"EnlargeProxyRequest", string(data)}, " ")
}
