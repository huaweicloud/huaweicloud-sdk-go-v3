package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeResponse Response Object
type ShowNodeResponse struct {

	// 设备ID
	Id *string `json:"id,omitempty"`

	// 设备架构
	Arch *string `json:"arch,omitempty"`

	// 设备内存
	Memory *int64 `json:"memory,omitempty"`

	// 设备名称，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64
	Name *string `json:"name,omitempty"`

	// 设备描述，最大长度255，不允许^, ~, ＃, $, %, &, *, <, >, (, ), [, ], {, }, ', \", \\
	Description *string `json:"description,omitempty"`

	// 产生时间，如2021-10-15 14:45:22 GMT+08:00
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间，如2021-10-15 14:45:22 GMT+08:00
	UpdatedAt *string `json:"updated_at,omitempty"`

	// IAM用户名
	UserName *string `json:"user_name,omitempty"`

	// 当该设备处于集群时，显示设备所属的集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 设备所处集群状态，集群创建（cluster_create）、集群删除（cluster_delete）、添加集群工作节点设备（cluster_add_nodes）、删除集群工作节点设备（cluster_delete_node）、集群节点设备状态更新（cluster_node_state_update）
	ClusterNodeState *string `json:"cluster_node_state,omitempty"`

	// 当该设备处于集群时，显示所属的集群设备类型。 - cluster_controller 控制设备 - cluster_worker 工作设备
	ClusterNodeType *string `json:"cluster_node_type,omitempty"`

	// 固件名称。可包含大小写字母、数字、下划线、中划线,长度不超过60字符。必须以字母开头,字母或数字结尾
	FirmwareName *string `json:"firmware_name,omitempty"`

	// 固件版本。支持X.Y.Z格式。每一个子版本号不超过三位且为非负整数,禁止在数字前补0
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	// 固件正在升级的版本
	UpgradeFirmwareVersion *string `json:"upgrade_firmware_version,omitempty"`

	// 固件升级状态，1、2、3分别代表升级中，升级失败，升级成
	FirmwareStatus *string `json:"firmware_status,omitempty"`

	FirmwareUpgradeRecord *[]FirmwareUpdateRecord `json:"firmware_upgrade_record,omitempty"`

	// 设备状态：UNCONNECTED(未注册)、RUNNING(运行中)、FAIL(故障)、STOPPED(停用)、UPGRADING(升级中)、FREEZE(冻结)
	State *string `json:"state,omitempty"`

	// 设备类型
	Type *string `json:"type,omitempty"`

	// 设备激活状态，未激活（INACTIVE）和已激活（ACTIVATED）
	ActiveStatus *string `json:"active_status,omitempty"`

	// 设备CPU个数
	Cpu *int32 `json:"cpu,omitempty"`

	// 设备GPU个数
	GpuNum *int32 `json:"gpu_num,omitempty"`

	// 设备NPU个数
	NpuNum *int32 `json:"npu_num,omitempty"`

	// 设备NPU类型
	NpuType *string `json:"npu_type,omitempty"`

	// 设备操作系统名称
	OsName *string `json:"os_name,omitempty"`

	// 设备操作系统类型
	OsType *string `json:"os_type,omitempty"`

	// 设备操作系统版本
	OsVersion *string `json:"os_version,omitempty"`

	// 是否启用容器
	EnableContainer *bool `json:"enable_container,omitempty"`

	// 是否启用GPU
	EnableGpu *bool `json:"enable_gpu,omitempty"`

	// 是否启用NPU
	EnableNpu *bool `json:"enable_npu,omitempty"`

	// 主机IP列表
	HostIps *[]string `json:"host_ips,omitempty"`

	Tags *NodeDetailResponseTags `json:"tags,omitempty"`

	// NPU信息
	NpuInfo *interface{} `json:"npu_info,omitempty"`

	// 激活订单号列表
	ActiveContent *[]string `json:"active_content,omitempty"`

	// 设备日志配置
	LogConfigs *[]LogConfig `json:"log_configs,omitempty"`

	// 事件有效时间(单位：分钟)
	EventValidityPeriod *int32 `json:"event_validity_period,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeResponse struct{}"
	}

	return strings.Join([]string{"ShowNodeResponse", string(data)}, " ")
}
