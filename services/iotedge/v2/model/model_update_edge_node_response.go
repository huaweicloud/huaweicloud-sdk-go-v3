package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeNodeResponse Response Object
type UpdateEdgeNodeResponse struct {

	// 边缘节点在IEF的日志配置
	LogConfigs *[]LogConfigDto `json:"log_configs,omitempty"`

	HaConfig *HaConfigDto `json:"ha_config,omitempty"`

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

	// AI加速卡类型，如华为昇腾AI加速卡NPU、图像处理加速卡GPU。
	AiCardType *string `json:"ai_card_type,omitempty"`

	// npu驱动动态库路径
	NpuLibraryPath *string `json:"npu_library_path,omitempty"`

	// 容器运行时版本
	ContainerVersion *string `json:"container_version,omitempty"`

	// 节点所属资源类型：advanced|standard
	Type *string `json:"type,omitempty"`

	// 节点的安全等级，MEDIUM边缘节数据上报不进行加密，HIGH对数据上报进行加密。
	SecurityLevel *string `json:"security_level,omitempty"`

	// 节点的可靠性等级。
	ReliabilityLevel *string `json:"reliability_level,omitempty"`

	// 节点的存储周期，默认0天，取值范围0~7天，0天则不存储。
	StoragePeriod *int32 `json:"storage_period,omitempty"`

	BasePath *BasePathDto `json:"base_path,omitempty"`

	// 注册节点网关配置
	HardwareModel *string `json:"hardware_model,omitempty"`

	OfflineCacheConfigs *OfflineCacheConfigsDto `json:"offline_cache_configs,omitempty"`

	DeviceAuthInfo *DeviceAuthInfoDisplayDto `json:"device_auth_info,omitempty"`

	// 节点使用的数据格式，默认为iotda物模型1.0格式，可以选择属性平铺数据格式flat_json
	DeviceDataFormat *string `json:"device_data_format,omitempty"`

	// 自动升级系统应用的节点开关，默认为关闭：OFF，IMMEDIATE表示节点开关打开
	AutomaticUpgrade *string `json:"automatic_upgrade,omitempty"`

	DeviceDataRecord *DeviceDataRecord `json:"device_data_record,omitempty"`

	// omagent监控运维工具是否上报指标
	MetricReport   *string `json:"metric_report,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeNodeResponse", string(data)}, " ")
}
