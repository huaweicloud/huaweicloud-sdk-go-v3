package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceStatusStatistics struct {

	// 统计对象名称(虚拟机组名称、桌面组名称)。
	StaticName *string `json:"static_name,omitempty"`

	// 运行数目(运行中、启动中、故障迁移中、迁移中)。
	RunningNum *int32 `json:"running_num,omitempty"`

	// 关机数目(关机中、已关机)。
	StopNum *int32 `json:"stop_num,omitempty"`

	// 休眠数目(休眠中、已休眠)。
	HibernateNum *int32 `json:"hibernate_num,omitempty"`

	// 其他(未知、删除失败、删除中)。
	OtherNum *int32 `json:"other_num,omitempty"`

	// 已分配数目
	AttachedNum *int32 `json:"attached_num,omitempty"`

	// 未分配数目
	UnattachedNum *int32 `json:"unattached_num,omitempty"`

	// 不可分配数目(分配失败、解分配失败、解分配成功)
	CannotAttachNum *int32 `json:"cannot_attach_num,omitempty"`

	// 处理中(分配中、解分配中)
	InProcessNum *int32 `json:"in_process_num,omitempty"`

	// 使用中数目。
	InUseNum *int32 `json:"in_use_num,omitempty"`

	// 未注册数目。
	UnregisteredNum *int32 `json:"unregistered_num,omitempty"`

	// 空闲数目。
	ReadyNum *int32 `json:"ready_num,omitempty"`

	// 断开连接数目。
	DisconnectedNum *int32 `json:"disconnected_num,omitempty"`

	// 未知数目。
	UnknownNum *int32 `json:"unknown_num,omitempty"`
}

func (o InstanceStatusStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceStatusStatistics struct{}"
	}

	return strings.Join([]string{"InstanceStatusStatistics", string(data)}, " ")
}
