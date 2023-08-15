package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceMetricRequestBody BCS组织实例监控数据详情查询请求结构
type ListInstanceMetricRequestBody struct {

	// 实体类型，[可选值如下: org(包括代理节点baas-agent、共识节点orderer、记账节点peer), plugin(插件)] 默认为org
	Type string `json:"type"`

	// 所属实体的名称，即区块链服务详情页面的“区块链实例”中各节点的名称（代理节点、共识节点、记账节点）
	EntityName string `json:"entity_name"`

	// 具体实例的名称。若当前区块链服务部署在CCE集群上，该名称为各节点对应的负载（pod）名称；若当前区块链服务部署在IEF集群上，可进入“智能边缘平台”服务，进入“边缘应用”->“容器应用”，查询负载名称。
	InstanceName string `json:"instance_name"`
}

func (o ListInstanceMetricRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMetricRequestBody struct{}"
	}

	return strings.Join([]string{"ListInstanceMetricRequestBody", string(data)}, " ")
}
