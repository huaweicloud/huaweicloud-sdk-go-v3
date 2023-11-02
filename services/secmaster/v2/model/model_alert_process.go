package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertProcess struct {

	// 进程名
	ProcessName *string `json:"process_name,omitempty"`

	// 进程执行文件路径
	ProcessPath *string `json:"process_path,omitempty"`

	// 进程id
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// 进程用户id
	ProcessUid *int32 `json:"process_uid,omitempty"`

	// 进程命令行
	ProcessCmdline *string `json:"process_cmdline,omitempty"`

	// 父进程名称
	ProcessParentName *string `json:"process_parent_name,omitempty"`

	// 父进程执行文件路径
	ProcessParentPath *string `json:"process_parent_path,omitempty"`

	// 父进程id
	ProcessParentPid *int32 `json:"process_parent_pid,omitempty"`

	// 父进程用户id
	ProcessParentUid *int32 `json:"process_parent_uid,omitempty"`

	// 父进程命令行
	ProcessParentCmdline *string `json:"process_parent_cmdline,omitempty"`

	// 子进程名称
	ProcessChildName *string `json:"process_child_name,omitempty"`

	// 子进程执行文件路径
	ProcessChildPath *string `json:"process_child_path,omitempty"`

	// 子进程id
	ProcessChildPid *int32 `json:"process_child_pid,omitempty"`

	// 子进程用户id
	ProcessChildUid *int32 `json:"process_child_uid,omitempty"`

	// 子进程命令行
	ProcessChildCmdline *string `json:"process_child_cmdline,omitempty"`

	// 进程启动时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	ProcessLauncheTime *string `json:"process_launche_time,omitempty"`

	// 进程结束时间，格式ISO8601：YYYY-MM-DDTHH:mm:ss.ms+timezone。时区信息为事件发生时区，无法解析时区的时间，默认时区填东八区
	ProcessTerminateTime *string `json:"process_terminate_time,omitempty"`
}

func (o AlertProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertProcess struct{}"
	}

	return strings.Join([]string{"AlertProcess", string(data)}, " ")
}
