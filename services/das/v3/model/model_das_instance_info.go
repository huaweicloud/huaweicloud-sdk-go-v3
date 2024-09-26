package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DasInstanceInfo struct {

	// 实例id。
	InstanceId string `json:"instance_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 实例状态。
	InstanceStatus string `json:"instance_status"`

	// 实例版本号。
	Version string `json:"version"`

	// 引擎类型。
	EngineType string `json:"engine_type"`

	// ip
	Ip string `json:"ip"`

	// 端口号
	Port int32 `json:"port"`

	// 实例cpu核数
	Cpu int32 `json:"cpu"`

	// 实例内存大小
	Mem int32 `json:"mem"`

	// 实例登录是否启用
	LoginFlag bool `json:"login_flag"`

	// 慢sql是否启用
	SlowSqlFlag bool `json:"slow_sql_flag"`

	// 死锁分析是否启用
	DeadlockFlag bool `json:"deadlock_flag"`

	// 锁阻塞是否启用
	LockBlockingFlag bool `json:"lock_blocking_flag"`

	// 当前实例是否计费
	ChargeFlag bool `json:"charge_flag"`

	// 全量sql是否启用
	FullSqlFlag bool `json:"full_sql_flag"`
}

func (o DasInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DasInstanceInfo struct{}"
	}

	return strings.Join([]string{"DasInstanceInfo", string(data)}, " ")
}
