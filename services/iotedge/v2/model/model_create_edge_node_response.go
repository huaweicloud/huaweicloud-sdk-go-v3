package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEdgeNodeResponse struct {

	// 边缘节点ID
	EdgeNodeId *string `json:"edge_node_id,omitempty" xml:"edge_node_id"`

	// 边缘节点名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 资源空间id，对应IOTDA云服务接口参数中的app_id。
	SpaceId *string `json:"space_id,omitempty" xml:"space_id"`

	// 边缘节点关联的产品ID，用于唯一标识一个产品模型。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`

	// 边缘节点关联的产品名称。
	ProductName *string `json:"product_name,omitempty" xml:"product_name"`

	// 边缘节点状态UNINSTALLED|INSTALLED|OFFLINE|ONLINE|DELETING|FROZEN
	State *string `json:"state,omitempty" xml:"state"`

	// 节点所属资源类型：advanced|standard
	Type *string `json:"type,omitempty" xml:"type"`

	// 安装文件版本
	InstallerVersion *string `json:"installer_version,omitempty" xml:"installer_version"`

	BasePath *BasePathDto `json:"base_path,omitempty" xml:"base_path"`

	// 资源id列表，创建节点时需绑定已购买的资源包，可以叠加节点功能。
	ResourceIds *[]string `json:"resource_ids,omitempty" xml:"resource_ids"`

	// 边缘节点ip列表
	Ips *[]string `json:"ips,omitempty" xml:"ips"`

	// 边缘节点创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 注册节点网关配置
	HardwareModel  *string `json:"hardware_model,omitempty" xml:"hardware_model"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeNodeResponse", string(data)}, " ")
}
