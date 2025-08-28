package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShrinkClusterReq struct {

	// 需要缩容的节点类型和数量集合。
	Shrink []ShrinkNodeReq `json:"shrink"`

	// 委托名称，委托给CSS服务，允许CSS调用您的其他云服务。
	AgencyName *string `json:"agency_name,omitempty"`

	// 操作类型。
	OperationType *string `json:"operation_type,omitempty"`

	// 是否需要检查集群负载。
	ClusterLoadCheck *bool `json:"cluster_load_check,omitempty"`
}

func (o ShrinkClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkClusterReq struct{}"
	}

	return strings.Join([]string{"ShrinkClusterReq", string(data)}, " ")
}
