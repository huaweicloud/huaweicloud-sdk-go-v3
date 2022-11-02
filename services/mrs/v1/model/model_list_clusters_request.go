package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClustersRequest struct {

	// 可以通过集群的标签来搜索指定标签的集群，当指定多个tag进行查询时，标签之间是与的关系。  - tags参数的格式为tags=k1*v1,k2*v2,k3*v3 - 当标签的value为空时，格式为tags=k1,k2,k3*v3
	Tags *string `json:"tags,omitempty"`

	// 分页查询每页返回的最大集群数量。  取值范围：[1～2147483646]
	PageSize *string `json:"pageSize,omitempty"`

	// 当前查询页码。
	CurrentPage *string `json:"currentPage,omitempty"`

	// 集群名称。
	ClusterName *string `json:"clusterName,omitempty"`

	// 根据集群状态查询集群列表。 - existing：查询现有集群列表，包括除“已删除”、包周期集群的“订单处理中”和“准备中”状态外的所有集群。 - history：查询历史集群列表，包括所有“已删除”、删除集群失败、集群删除虚拟机失败、删除集群更新数据库失败等状态的集群。 - starting：查询启动中的集群列表。 - running：查询运行中的集群列表。 - terminated：查询已删除的集群列表。 - failed：查询失败的集群列表。 - abnormal：查询异常的集群列表。 - terminating：查询删除中的集群列表。 - frozen：查询已冻结的集群列表。 - scaling-out：查询扩容中的集群列表。 - scaling-in：查询缩容中的集群列表。
	ClusterState *string `json:"clusterState,omitempty"`

	// 通过企业项目ID来搜索指定项目的集群。  该参数默认设置为0，表示为default企业项目。  获取方式请参见《企业管理API参考》的“查询企业项目列表”响应消息表“enterprise_project字段数据结构说明”的“id”。
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`
}

func (o ListClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}
