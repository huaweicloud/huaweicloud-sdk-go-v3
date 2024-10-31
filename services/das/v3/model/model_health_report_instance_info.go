package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportInstanceInfo struct {

	// 账号ID。
	TenantId string `json:"tenant_id"`

	// 项目ID。
	ProjectId string `json:"project_id"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 主节点ID。
	MasterNodeId string `json:"master_node_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 实例cpu核数。
	Cpu int32 `json:"cpu"`

	// 实例内存大小。
	Mem int32 `json:"mem"`

	// 磁盘大小。
	DiskSize int32 `json:"disk_size"`

	// 磁盘类型。
	DiskType string `json:"disk_type"`

	// 实例引擎类型。
	Engine string `json:"engine"`

	// 引擎内核版本。
	EngineVersion string `json:"engine_version"`
}

func (o HealthReportInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportInstanceInfo struct{}"
	}

	return strings.Join([]string{"HealthReportInstanceInfo", string(data)}, " ")
}
