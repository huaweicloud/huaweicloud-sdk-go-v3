package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterNode struct {

	// 微服务引擎CCE节点ID
	Id *string `json:"id,omitempty"`

	// 微服务引擎CCE节点所属可用区
	Az *string `json:"az,omitempty"`

	// 微服务引擎CCE节点IP
	Ip *string `json:"ip,omitempty"`

	// 微服务引擎CCE节点标签
	Label *string `json:"label,omitempty"`

	// 微服务引擎CCE节点状态
	Status *string `json:"status,omitempty"`
}

func (o ClusterNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNode struct{}"
	}

	return strings.Join([]string{"ClusterNode", string(data)}, " ")
}
