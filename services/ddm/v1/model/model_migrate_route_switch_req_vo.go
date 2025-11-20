package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateRouteSwitchReqVo struct {
	IamAccount *IamAccount `json:"iam_account,omitempty"`

	// 项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务id。
	TaskId *string `json:"task_id,omitempty"`

	// 自动切换路由开始时间。
	SwitchRouteBeginTime *string `json:"switch_route_begin_time,omitempty"`

	// 自动切换路由结束时间。
	SwitchRouteEndTime *string `json:"switch_route_end_time,omitempty"`

	// 是否openapi。
	IsOpenApi *bool `json:"is_open_api,omitempty"`

	// 逻辑库名称。
	LogicDbName *string `json:"logic_db_name,omitempty"`
}

func (o MigrateRouteSwitchReqVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateRouteSwitchReqVo struct{}"
	}

	return strings.Join([]string{"MigrateRouteSwitchReqVo", string(data)}, " ")
}
