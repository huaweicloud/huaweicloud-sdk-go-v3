package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例数量统计信息。
type StatusStatistic struct {

	// 支付中的实例数。
	PayingCount *int32 `json:"paying_count,omitempty" xml:"paying_count"`

	// 冻结中的实例数。
	FreezingCount *int32 `json:"freezing_count,omitempty" xml:"freezing_count"`

	// 迁移中的实例数。
	MigratingCount *int32 `json:"migrating_count,omitempty" xml:"migrating_count"`

	// 清空中的实例数。
	FlushingCount *int32 `json:"flushing_count,omitempty" xml:"flushing_count"`

	// 升级中的实例数。
	UpgradingCount *int32 `json:"upgrading_count,omitempty" xml:"upgrading_count"`

	// 恢复中的实例数。
	RestoringCount *int32 `json:"restoring_count,omitempty" xml:"restoring_count"`

	// 扩容中的实例数。
	ExtendingCount *int32 `json:"extending_count,omitempty" xml:"extending_count"`

	// 正在创建的实例数。
	CreatingCount *int32 `json:"creating_count,omitempty" xml:"creating_count"`

	// 正在运行的实例数。
	RunningCount *int32 `json:"running_count,omitempty" xml:"running_count"`

	// 异常的实例数。
	ErrorCount *int32 `json:"error_count,omitempty" xml:"error_count"`

	// 已冻结的实例数。
	FrozenCount *int32 `json:"frozen_count,omitempty" xml:"frozen_count"`

	// 创建失败的实例数。
	CreatefailedCount *int32 `json:"createfailed_count,omitempty" xml:"createfailed_count"`

	// 正在重启的实例数。
	RestartingCount *int32 `json:"restarting_count,omitempty" xml:"restarting_count"`
}

func (o StatusStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusStatistic struct{}"
	}

	return strings.Join([]string{"StatusStatistic", string(data)}, " ")
}
