package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiagnosisTask 诊断记录对象
type DiagnosisTask struct {

	// id
	Id *string `json:"id,omitempty"`

	// code
	Code *string `json:"code,omitempty"`

	// 诊断记录所属租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 被诊断实例所属项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 诊断记录所属用户ID
	UserId *string `json:"user_id,omitempty"`

	// 诊断记录所属用户名称
	UserName *string `json:"user_name,omitempty"`

	// 诊断进度
	Progress *int32 `json:"progress,omitempty"`

	// 诊断任务工单ID
	WorkOrderId *string `json:"work_order_id,omitempty"`

	// 被诊断的实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 被诊断的实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 被诊断实例类别，ECS、RDS等
	Type *string `json:"type,omitempty"`

	// 诊断任务状态执行状态
	Status *string `json:"status,omitempty"`

	// 创建时间，遵循RFC3339规范，精确到秒 示例：2020-09-01T18:50:20Z
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，遵循RFC3339规范，精确到秒 示例：2020-09-01T18:50:20Z
	EndTime *string `json:"end_time,omitempty"`

	// 被诊断实例的数量
	InstanceNum *int32 `json:"instance_num,omitempty"`

	// 被诊断实例的操作系统类型
	OsType *string `json:"os_type,omitempty"`

	// 被诊断实例所在区域
	Region *string `json:"region,omitempty"`
}

func (o DiagnosisTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisTask struct{}"
	}

	return strings.Join([]string{"DiagnosisTask", string(data)}, " ")
}
