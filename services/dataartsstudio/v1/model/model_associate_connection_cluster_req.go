package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateConnectionClusterReq 绑定集群的请求body体。
type AssociateConnectionClusterReq struct {

	// 需要使用资源组网络连接的集群名列表。单条最大长度128字符。
	Clusters []string `json:"clusters"`
}

func (o AssociateConnectionClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateConnectionClusterReq struct{}"
	}

	return strings.Join([]string{"AssociateConnectionClusterReq", string(data)}, " ")
}
