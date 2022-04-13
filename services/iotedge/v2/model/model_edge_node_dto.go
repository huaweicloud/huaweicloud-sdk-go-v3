package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询边缘节点列表响应体
type EdgeNodeDto struct {
	// 边缘节点ID

	EdgeNodeId *string `json:"edge_node_id,omitempty"`
	// 边缘节点名称

	Name *string `json:"name,omitempty"`
	// 边缘节点状态UNINSTALLED|INSTALLED|OFFLINE|ONLINE|DELETING|UPGRADING

	State *string `json:"state,omitempty"`
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"instance_id,omitempty"`
	// 资源空间id，对应IOTDA云服务接口参数中的app_id。

	SpaceId *string `json:"space_id,omitempty"`
	// 节点所属资源类型：advanced|standard

	Type *string `json:"type,omitempty"`
	// 节点所购买的资源类型的列表

	ResourceIds *[]string `json:"resource_ids,omitempty"`
	// 节点所购买的资源类型的列表

	ResourceSpecTypes *[]string `json:"resource_spec_types,omitempty"`
	// 边缘节点ip列表

	Ips *[]string `json:"ips,omitempty"`
	// 边缘节点创建时间

	CreateTime *string `json:"create_time,omitempty"`
}

func (o EdgeNodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeDto struct{}"
	}

	return strings.Join([]string{"EdgeNodeDto", string(data)}, " ")
}
