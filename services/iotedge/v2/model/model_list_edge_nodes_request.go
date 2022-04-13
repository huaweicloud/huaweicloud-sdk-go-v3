package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEdgeNodesRequest struct {
	// 节点名称

	Name *string `json:"name,omitempty"`
	// 节点状态,OFFLINE|ONLINE|UNINSTALLED|INSTALLED|DELETING|UPGRADING

	State *string `json:"state,omitempty"`
	// 节点所属资源类型，advanced|standard

	Type *string `json:"type,omitempty"`
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"instance_id,omitempty"`
	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，可以携带该参数查询指定资源空间下的设备列表，不携带该参数则会查询该用户下所有设备列表。

	SpaceId *string `json:"space_id,omitempty"`
	// 节点id列表,查询ID在给的节点ID列表内的节点信息

	NodeIds *[]string `json:"node_ids,omitempty"`
	// 查询的起始位置，取值范围为非负整数，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数，默认值为10，取值区间为1-1000

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEdgeNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeNodesRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeNodesRequest", string(data)}, " ")
}
