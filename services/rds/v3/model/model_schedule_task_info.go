package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleTaskInfo struct {

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

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 任务创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务计划开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartTime *string `json:"start_time,omitempty"`

	// 任务计划结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime *string `json:"end_time,omitempty"`

	// 任务优先级，整数取值范围1-100，数值越小优先级越高。
	Order *string `json:"order,omitempty"`

	// 任务状态。取值范围如下： Initing:初始化中 Pending:挂起 Running:运行中 Completed:已完成 Failed:已失败 Unauthorized:未授权 Canceled:已取消
	Status *string `json:"status,omitempty"`

	// 数据库类型。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 磁盘类型。
	VolumeType *string `json:"volume_type,omitempty"`

	TargetConfig *TargetConfig `json:"target_config,omitempty"`
}

func (o ScheduleTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleTaskInfo struct{}"
	}

	return strings.Join([]string{"ScheduleTaskInfo", string(data)}, " ")
}
