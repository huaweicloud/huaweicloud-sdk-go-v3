package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEdgeNodeResponse struct {

	// 边缘节点在IEF的日志配置
	LogConfigs *[]LogConfigDto `json:"log_configs,omitempty"`

	// 边缘节点Id
	EdgeNodeId *string `json:"edge_node_id,omitempty"`

	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。
	InstanceId *string `json:"instance_id,omitempty"`

	// 边缘节点关联的产品ID，用于唯一标识一个产品模型。
	ProductId *string `json:"product_id,omitempty"`

	// 边缘节点关联的产品名称。
	ProductName *string `json:"product_name,omitempty"`

	// 资源空间id，对应IOTDA云服务接口参数中的app_id。
	SpaceId *string `json:"space_id,omitempty"`

	// 节点所购买的资源类型的列表
	ResourceSpecTypes *[]string `json:"resource_spec_types,omitempty"`

	// 资源id列表，创建节点时需绑定已购买的资源包，可以叠加节点功能。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 边缘节点主机ip
	Ips *[]string `json:"ips,omitempty"`

	// 边缘节点名称
	Name *string `json:"name,omitempty"`

	// 边缘节点状态
	State *string `json:"state,omitempty"`

	// 边缘应用id，只允许数字、英文小写、中划线，切必须以字母或数字结尾
	SoftwareVersion *string `json:"software_version,omitempty"`

	// 边缘节点创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 边缘节点更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 边缘节点操作系统名称
	OsName *string `json:"os_name,omitempty"`

	// 边缘节点操作系统架构
	Arch *string `json:"arch,omitempty"`

	// 边缘节点主机名
	HostName *string `json:"host_name,omitempty"`

	// 边缘节点网络网卡信息
	Nics *[]Nic `json:"nics,omitempty"`

	// 网络规格，如4 cores | 3867 MB
	Specification *string `json:"specification,omitempty"`

	// 华为AI加速卡类型，如NPU、GPU、unEquipped
	AiCardType *string `json:"ai_card_type,omitempty"`

	// 容器运行时版本
	ContainerVersion *string `json:"container_version,omitempty"`

	// 节点所属资源类型：advanced|standard
	Type *string `json:"type,omitempty"`

	// 节点的安全等级，MEDIUM边缘节数据上报不进行加密，HIGH对数据上报进行加密。
	SecurityLevel *string `json:"security_level,omitempty"`

	// 节点的存储周期，默认0天，取值范围0~7天，0天则不存储。
	StoragePeriod *int32 `json:"storage_period,omitempty"`

	BasePath       *BasePathDto `json:"base_path,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeResponse", string(data)}, " ")
}
