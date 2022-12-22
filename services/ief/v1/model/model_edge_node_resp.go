package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点参数
type EdgeNodeResp struct {

	// 边缘节点ID
	Id string `json:"id"`

	// 边缘节点名称，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64 Name为必填字段，且本帐号中唯一。
	Name string `json:"name"`

	// 边缘节点描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description string `json:"description"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 私钥
	PrivateKey string `json:"private_key"`

	// 证书
	Certificate string `json:"certificate"`

	// 根证书
	Ca string `json:"ca"`

	// 边缘节点状态 - UNCONNECTED（未注册） - RUNNING（运行中） - FAIL（故障） - STOPPED（停用） - UPGRADING（升级中） - FREEZE（冻结）
	State string `json:"state"`

	// 将证书文件certificate/ca/private_key打成.tar.gz包后用base64编码的字符串。 使用时请使用base64解码成.tar.gz包。
	Package string `json:"package"`

	// 云端服务URL
	MasterAddr string `json:"master_addr"`

	// 边缘节点CPU核心数
	Cpu int32 `json:"cpu"`

	// 边缘节点内存大小，单位M
	Memory int32 `json:"memory"`

	// 边缘节点操作系统名称
	OsName string `json:"os_name"`

	// 边缘节点操作系统版本
	OsVersion string `json:"os_version"`

	// pause容器镜像URL
	PauseDockerImage string `json:"pause_docker_image"`

	// 边缘节点架构
	Arch string `json:"arch"`

	// 边缘节点操作系统类型
	OsType string `json:"os_type"`

	// 部署在该边缘节点上的应用实例个数
	DeploymentNum int32 `json:"deployment_num"`

	// 边缘节点是否开启GPU，默认为false
	EnableGpu bool `json:"enable_gpu"`

	// 边缘节点日志配置
	LogConfigs []LogConfigs `json:"log_configs"`

	// 关联设备信息
	DeviceInfos []DeviceInfos `json:"device_infos"`

	// edged版本
	EdgedVersion string `json:"edged_version"`

	// gpu个数
	GpuNum int32 `json:"gpu_num"`

	// 主机IP，默认返回eth0网卡的IP。
	HostIps []string `json:"host_ips"`

	// 与device绑定关系名称（通过device id查询node时有值）
	Relation string `json:"relation"`

	// 与device绑定关系描述（通过device id查询node时有值）
	Comment string `json:"comment"`

	// gpu型号和gpu memory大小
	GpuInfo []GpuInfo `json:"gpu_info"`

	// 关联设备数量
	DeviceNum int32 `json:"device_num"`

	// 边缘节点是否开启NPU，true表示开启，false表示不开启，默认为false
	EnableNpu bool `json:"enable_npu"`

	// NPU类型，支持D310类型和D910类型。 - D310表示D310类型。 - D910表示D910类型。 - 不填表示为D310类型。
	NpuType string `json:"npu_type"`

	// 节点网卡和对应IP地址信息
	Nics []Nics `json:"nics"`

	// 边缘节点主机名
	HostName string `json:"host_name"`

	// 边缘节点版本
	IefNodeVersion string `json:"ief_node_version"`

	// 是否能升级的标志 - true：需要升级 - false：不需要升级
	UpgradeFlag bool `json:"upgrade_flag"`

	// 产品ID（通过产品证书方式纳管）
	ProductId string `json:"product_id"`

	// 节点组ID（一个节点属于多个节点组）
	GroupIds []string `json:"group_ids"`

	// 节点安装或升级记录
	UpgradeHistory []UpgradeHistory `json:"upgrade_history"`

	// 边缘节点属性，关联属性个数最多为32个
	Attributes []Attributes `json:"attributes"`

	// 节点是否开启Docker
	DockerEnable bool `json:"docker_enable"`

	// mqtt集成模式 - internal：edgecore内置mqtt - external：外置开源mqtt
	MqttMode string `json:"mqtt_mode"`

	// 外置开源mqtt地址
	MqttExternal string `json:"mqtt_external"`

	// edgecore内置的mqtt地址
	MqttInternal string `json:"mqtt_internal"`

	// 节点类型，默认为空，非空时为小站节点
	NodeType string `json:"node_type"`

	NtpConfigs *NtpConfigs `json:"ntp_configs"`

	// 节点故障原因
	ErrorReason string `json:"error_reason"`

	// 边缘节点标签，标签个数最多为20个
	Tags []ResourceTag `json:"tags"`

	// NPU数量
	NpuNum int32 `json:"npu_num"`

	// NPU型号和NPU Memory大小
	NpuInfo []NpuInfo `json:"npu_info"`

	// 容器运行时版本
	ContainerRuntimeVersion string `json:"container_runtime_version"`

	// 边缘节点使用token注册时的凭证
	Identifier *string `json:"identifier,omitempty"`

	// IEC/IES节点id
	PurchaseId *string `json:"purchase_id,omitempty"`

	StateDetails *StateDetails `json:"state_details,omitempty"`

	// 证书有效期持续时间
	CertRemainingValidTime *int32 `json:"cert_remaining_valid_time,omitempty"`
}

func (o EdgeNodeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeResp struct{}"
	}

	return strings.Join([]string{"EdgeNodeResp", string(data)}, " ")
}
