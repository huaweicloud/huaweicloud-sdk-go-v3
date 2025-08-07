package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Task struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 任务名称。
	Name *string `json:"name,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态。 - 值为“BUILD”，表示实例正在创建。 - 值为“CREATE FAIL”，表示实例创建失败。 - 值为“ACTIVE”，表示实例正常。 - 值为“FAILED”，表示实例异常。 - 值为“FROZEN”，表示实例冻结。 - 值为“MODIFYING”，表示实例正在扩容。 - 值为“REBOOTING”，表示实例正在重启。 - 值为“RESTORING”，表示实例正在恢复。 - 值为“MODIFYING INSTANCE TYPE”，表示实例正在转主备。 - 值为“SWITCHOVER”，表示实例正在主备切换。 - 值为“MIGRATING”，表示实例正在迁移。 - 值为“BACKING UP”，表示实例正在进行备份。 - 值为“MODIFYING DATABASE PORT”，表示实例正在修改数据库端口。 - 值为“STORAGE FULL”，表示实例磁盘空间满。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 任务进度，单位百分比，取值范围：0-100。当任务成功或者失败状态下，该值为空字符串。
	Process *string `json:"process,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`

	// 任务创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务计划结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态。取值范围如下： Running:运行中 Completed:已完成 Failed:已失败
	Status *string `json:"status,omitempty"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}
