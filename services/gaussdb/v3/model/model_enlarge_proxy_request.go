package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// proxy节点扩容信息
type EnlargeProxyRequest struct {

	// proxy节点扩容操作需要扩容的节点数。  扩容的节点数的取值范围：1~30之间的整数。  限制条件：该实例的proxy节点的总数量小于等于32。
	NodeNum int32 `json:"node_num"`

	// 数据库代理id。  如果实例只开启了一个代理，可不传该参数；如果实例开启了多个代理，则必须指定一个数据库代理，扩容新的代理节点。
	ProxyId *string `json:"proxy_id,omitempty"`
}

func (o EnlargeProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeProxyRequest struct{}"
	}

	return strings.Join([]string{"EnlargeProxyRequest", string(data)}, " ")
}
