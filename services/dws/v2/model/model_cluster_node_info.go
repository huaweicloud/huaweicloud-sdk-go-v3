package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterNodeInfo 逻辑集群节点信息
type ClusterNodeInfo struct {

	// 节点ID
	Id *string `json:"id,omitempty"`

	// 节点名称
	Name *string `json:"name,omitempty"`

	// 节点状态
	Status *string `json:"status,omitempty"`

	// 节点子状态
	SubStatus *string `json:"sub_status,omitempty"`

	// 节点规格
	Spec *string `json:"spec,omitempty"`

	// 实例创建类型
	InstCreateType *string `json:"inst_create_type,omitempty"`

	// 节点别名
	AliasName *string `json:"alias_name,omitempty"`

	// 可用区编码
	AzCode *string `json:"az_code,omitempty"`
}

func (o ClusterNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNodeInfo struct{}"
	}

	return strings.Join([]string{"ClusterNodeInfo", string(data)}, " ")
}
