package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventResponseExecuteOptions struct {

	// 本地盘设备名
	Device *string `json:"device,omitempty"`

	// 本地盘挂载唯一标识
	Wwn *string `json:"wwn,omitempty"`

	// 本地盘序列号
	SerialNumber *string `json:"serial_number,omitempty"`

	// flavorID
	ResizeTargetFlavorId *string `json:"resize_target_flavor_id,omitempty"`

	// 实例迁移策略
	MigratePolicy *string `json:"migrate_policy,omitempty"`

	// 执行器
	Executor *string `json:"executor,omitempty"`
}

func (o EventResponseExecuteOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResponseExecuteOptions struct{}"
	}

	return strings.Join([]string{"EventResponseExecuteOptions", string(data)}, " ")
}
