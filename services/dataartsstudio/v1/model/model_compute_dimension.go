package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputeDimension struct {

	// 节点实例ID
	NodeInstanceId *string `json:"node_instance_id,omitempty"`

	// 节点名称或脚本名称
	Name *string `json:"name,omitempty"`

	// 类型,用于区分节点实例和脚本
	Type *int32 `json:"type,omitempty"`

	// 节点类型
	NodeType *string `json:"node_type,omitempty"`

	// 作业id
	JobId *int64 `json:"job_id,omitempty"`

	// 作业实例id
	JobInstanceId *int64 `json:"job_instance_id,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 节点实例启动时间,单位毫秒
	StartTimeMs *int64 `json:"start_time_ms,omitempty"`

	// 节点实例结束时间,单位毫秒
	EndTimeMs *int64 `json:"end_time_ms,omitempty"`

	// 节点实例执行时长,单位毫秒
	ExecuteTimeMs *int64 `json:"execute_time_ms,omitempty"`

	// 节点实例状态
	Status *string `json:"status,omitempty"`

	// 节点实例上次执行开始时间,单位毫秒
	LastStartTimeMs *int64 `json:"last_start_time_ms,omitempty"`

	// 节点实例上次执行结束时间,单位毫秒
	LastEndTimeMs *int64 `json:"last_end_time_ms,omitempty"`

	// cpu消耗值
	CpuResourceConsumptionValue *int64 `json:"cpu_resource_consumption_value,omitempty"`

	// cpu消耗单位
	CpuUnit *string `json:"cpu_unit,omitempty"`

	// 内存消耗值
	MemoryResourceConsumptionValue *int64 `json:"memory_resource_consumption_value,omitempty"`

	// 内存消耗单位
	MemoryUnit *string `json:"memory_unit,omitempty"`

	// 上次cpu消耗值
	LastCpuResourceConsumptionValue *int64 `json:"last_cpu_resource_consumption_value,omitempty"`

	// 上次cpu消耗单位
	LastCpuUnit *string `json:"last_cpu_unit,omitempty"`

	// 上次内存消耗值
	LastMemoryResourceConsumptionValue *int64 `json:"last_memory_resource_consumption_value,omitempty"`

	// 上次内存消耗单位
	LastMemoryUnit *string `json:"last_memory_unit,omitempty"`

	// 节点实例创建者name
	Creator *string `json:"creator,omitempty"`

	// 数据入库时间,单位毫秒
	WarehouseTimeMs *int64 `json:"warehouse_time_ms,omitempty"`

	// 扩展字段
	ExtendedFields *string `json:"extended_fields,omitempty"`
}

func (o ComputeDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeDimension struct{}"
	}

	return strings.Join([]string{"ComputeDimension", string(data)}, " ")
}
