package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEdgeNodeResponse struct {
	// 边缘节点ID

	EdgeNodeId *string `json:"edge_node_id,omitempty"`
	// 边缘节点名称

	Name *string `json:"name,omitempty"`
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"instance_id,omitempty"`
	// 资源空间id，对应IOTDA云服务接口参数中的app_id。

	SpaceId *string `json:"space_id,omitempty"`
	// 边缘节点关联的产品ID，用于唯一标识一个产品模型。

	ProductId *string `json:"product_id,omitempty"`
	// 边缘节点关联的产品名称。

	ProductName *string `json:"product_name,omitempty"`
	// 边缘节点状态UNINSTALLED|INSTALLED|OFFLINE|ONLINE|DELETING|UPGRADING

	State *string `json:"state,omitempty"`
	// 节点所属资源类型：advanced|standard

	Type *string `json:"type,omitempty"`
	// 安装文件版本

	InstallerVersion *string `json:"installer_version,omitempty"`
	// 资源id列表，创建节点时需绑定已购买的资源包，可以叠加节点功能。

	ResourceIds *[]string `json:"resource_ids,omitempty"`
	// 边缘节点ip列表

	Ips *[]string `json:"ips,omitempty"`
	// 边缘节点创建时间

	CreateTime     *string `json:"create_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeNodeResponse", string(data)}, " ")
}
