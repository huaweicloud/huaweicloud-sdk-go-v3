package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设备详情
type NodeResponse struct {

	// 设备ID
	Id *string `json:"id,omitempty"`

	// 设备名称，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64
	Name *string `json:"name,omitempty"`

	// 设备描述，最大长度255，不允许^, ~, ＃, $, %, &, *, <, >, (, ), [, ], {, }, ', \", \\
	Description *string `json:"description,omitempty"`

	// 产生时间，如2021-10-15 14:45:22 GMT+08:00
	CreatedAt *string `json:"created_at,omitempty"`

	// 当该设备处于集群时，显示设备所属的集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 设备所处集群状态，集群创建（cluster_create）、集群删除（cluster_delete）、添加集群工作节点设备（cluster_add_nodes）、删除集群工作节点设备（cluster_delete_node）、集群节点设备状态更新（cluster_node_state_update）
	ClusterNodeState *string `json:"cluster_node_state,omitempty"`

	// 当该设备处于集群时，显示所属的集群设备类型。 - cluster_controller 控制设备 - cluster_worker 工作设备
	ClusterNodeType *string `json:"cluster_node_type,omitempty"`

	// 固件名称。可包含大小写字母、数字、下划线、中划线,长度不超过60字符。必须以字母开头,字母或数字结尾
	FirmwareName *string `json:"firmware_name,omitempty"`

	// 固件正在升级的版本
	UpgradeFirmwareVersion *string `json:"upgrade_firmware_version,omitempty"`

	// 固件升级状态，1、2、3分别代表升级中，升级失败，升级成功
	FirmwareStatus *int32 `json:"firmware_status,omitempty"`

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

	// 主机IP列表
	HostIps *[]string `json:"host_ips,omitempty"`
}

func (o NodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeResponse struct{}"
	}

	return strings.Join([]string{"NodeResponse", string(data)}, " ")
}
