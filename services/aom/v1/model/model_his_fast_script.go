package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HisFastScript HIS 用户快速执行脚本实体类。
type HisFastScript struct {

	// 脚本类型。SHELL BAT PYTHON POWER_SHELL
	ScriptType string `json:"script_type"`

	// 执行脚本的ECS机器用户。
	CmdUser string `json:"cmd_user"`

	// 脚本内容。
	ScriptContent string `json:"script_content"`

	// 执行的机器列表。
	EcsIdList string `json:"ecs_id_list"`

	// 任务名称。
	Name string `json:"name"`

	// 项目ID。
	ProjectId string `json:"project_id"`

	// 任务参数，多个参数以空格分隔。
	ScriptArgs *string `json:"script_args,omitempty"`
}

func (o HisFastScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HisFastScript struct{}"
	}

	return strings.Join([]string{"HisFastScript", string(data)}, " ")
}
