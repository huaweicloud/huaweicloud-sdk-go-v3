package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 源端列表中关联的任务
type TaskByServerSource struct {

	// 任务id
	Id *string `json:"id,omitempty" xml:"id"`

	// 任务名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 任务类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 任务状态
	State *string `json:"state,omitempty" xml:"state"`

	// 开始时间
	StartDate *int64 `json:"start_date,omitempty" xml:"start_date"`

	// 限速
	SpeedLimit *int32 `json:"speed_limit,omitempty" xml:"speed_limit"`

	// 迁移速率
	MigrateSpeed *float64 `json:"migrate_speed,omitempty" xml:"migrate_speed"`

	// 是否启动虚拟机
	StartTargetServer *bool `json:"start_target_server,omitempty" xml:"start_target_server"`

	// 虚拟机模板id
	VmTemplateId *string `json:"vm_template_id,omitempty" xml:"vm_template_id"`

	// region_id
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	TargetServer *TargetServerById `json:"target_server,omitempty" xml:"target_server"`

	// 日志收集状态
	LogCollectStatus *string `json:"log_collect_status,omitempty" xml:"log_collect_status"`

	// 是否使用已有虚拟机
	ExistServer *bool `json:"exist_server,omitempty" xml:"exist_server"`

	// 是否使用公网ip
	UsePublicIp *bool `json:"use_public_ip,omitempty" xml:"use_public_ip"`

	CloneServer *CloneServer `json:"clone_server,omitempty" xml:"clone_server"`
}

func (o TaskByServerSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskByServerSource struct{}"
	}

	return strings.Join([]string{"TaskByServerSource", string(data)}, " ")
}
