package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DasInstanceInfoDto DAS实例信息
type DasInstanceInfoDto struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 实例版本号
	Version *string `json:"version,omitempty"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 客户端IP
	Ip *string `json:"ip,omitempty"`

	// 客户端端口号
	Port *int32 `json:"port,omitempty"`

	// 实例cpu核数
	Cpu *int32 `json:"cpu,omitempty"`

	// 实例内存大小
	Mem *int32 `json:"mem,omitempty"`

	// 实例登录是否启用
	LoginFlag *bool `json:"login_flag,omitempty"`

	// 慢sql是否启用
	SlowSqlFlag *bool `json:"slow_sql_flag,omitempty"`

	// 死锁分析是否启用
	DeadLockFlag *bool `json:"dead_lock_flag,omitempty"`

	// 锁阻塞是否启用
	LockBlockingFlag *bool `json:"lock_blocking_flag,omitempty"`

	// 当前实例是否计费
	ChargeFlag *bool `json:"charge_flag,omitempty"`

	// 实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 全量sql是否启用
	FullSqlFlag *bool `json:"full_sql_flag,omitempty"`
}

func (o DasInstanceInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DasInstanceInfoDto struct{}"
	}

	return strings.Join([]string{"DasInstanceInfoDto", string(data)}, " ")
}
