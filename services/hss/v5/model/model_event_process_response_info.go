package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 进程信息
type EventProcessResponseInfo struct {

	// 进程名称
	ProcessName *string `json:"process_name,omitempty" xml:"process_name"`

	// 进程文件路径
	ProcessPath *string `json:"process_path,omitempty" xml:"process_path"`

	// 进程id
	ProcessPid *int32 `json:"process_pid,omitempty" xml:"process_pid"`

	// 进程用户id
	ProcessUid *int32 `json:"process_uid,omitempty" xml:"process_uid"`

	// 运行进程的用户名
	ProcessUsername *string `json:"process_username,omitempty" xml:"process_username"`

	// 进程文件命令行
	ProcessCmdline *string `json:"process_cmdline,omitempty" xml:"process_cmdline"`

	// 进程文件名
	ProcessFilename *string `json:"process_filename,omitempty" xml:"process_filename"`

	// 进程启动时间
	ProcessStartTime *int64 `json:"process_start_time,omitempty" xml:"process_start_time"`

	// 进程组ID
	ProcessGid *int32 `json:"process_gid,omitempty" xml:"process_gid"`

	// 进程有效组ID
	ProcessEgid *int32 `json:"process_egid,omitempty" xml:"process_egid"`

	// 进程有效用户ID
	ProcessEuid *int32 `json:"process_euid,omitempty" xml:"process_euid"`

	// 父进程名称
	ParentProcessName *string `json:"parent_process_name,omitempty" xml:"parent_process_name"`

	// 父进程文件路径
	ParentProcessPath *string `json:"parent_process_path,omitempty" xml:"parent_process_path"`

	// 父进程id
	ParentProcessPid *int32 `json:"parent_process_pid,omitempty" xml:"parent_process_pid"`

	// 父进程用户id
	ParentProcessUid *int32 `json:"parent_process_uid,omitempty" xml:"parent_process_uid"`

	// 父进程文件命令行
	ParentProcessCmdline *string `json:"parent_process_cmdline,omitempty" xml:"parent_process_cmdline"`

	// 父进程文件名
	ParentProcessFilename *string `json:"parent_process_filename,omitempty" xml:"parent_process_filename"`

	// 父进程启动时间
	ParentProcessStartTime *int64 `json:"parent_process_start_time,omitempty" xml:"parent_process_start_time"`

	// 父进程组ID
	ParentProcessGid *int32 `json:"parent_process_gid,omitempty" xml:"parent_process_gid"`

	// 父进程有效组ID
	ParentProcessEgid *int32 `json:"parent_process_egid,omitempty" xml:"parent_process_egid"`

	// 父进程有效用户ID
	ParentProcessEuid *int32 `json:"parent_process_euid,omitempty" xml:"parent_process_euid"`

	// 子进程名称
	ChildProcessName *string `json:"child_process_name,omitempty" xml:"child_process_name"`

	// 子进程文件路径
	ChildProcessPath *string `json:"child_process_path,omitempty" xml:"child_process_path"`

	// 子进程id
	ChildProcessPid *int32 `json:"child_process_pid,omitempty" xml:"child_process_pid"`

	// 子进程用户id
	ChildProcessUid *int32 `json:"child_process_uid,omitempty" xml:"child_process_uid"`

	// 子进程文件命令行
	ChildProcessCmdline *string `json:"child_process_cmdline,omitempty" xml:"child_process_cmdline"`

	// 子进程文件名
	ChildProcessFilename *string `json:"child_process_filename,omitempty" xml:"child_process_filename"`

	// 子进程启动时间
	ChildProcessStartTime *int64 `json:"child_process_start_time,omitempty" xml:"child_process_start_time"`

	// 子进程组ID
	ChildProcessGid *int32 `json:"child_process_gid,omitempty" xml:"child_process_gid"`

	// 子进程有效组ID
	ChildProcessEgid *int32 `json:"child_process_egid,omitempty" xml:"child_process_egid"`

	// 子进程有效用户ID
	ChildProcessEuid *int32 `json:"child_process_euid,omitempty" xml:"child_process_euid"`

	// 虚拟化命令
	VirtCmd *string `json:"virt_cmd,omitempty" xml:"virt_cmd"`

	// 虚拟化进程名称
	VirtProcessName *string `json:"virt_process_name,omitempty" xml:"virt_process_name"`

	// 逃逸方式
	EscapeMode *string `json:"escape_mode,omitempty" xml:"escape_mode"`

	// 逃逸后后执行的命令
	EscapeCmd *string `json:"escape_cmd,omitempty" xml:"escape_cmd"`

	// 进程启动文件hash
	ProcessHash *string `json:"process_hash,omitempty" xml:"process_hash"`
}

func (o EventProcessResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventProcessResponseInfo struct{}"
	}

	return strings.Join([]string{"EventProcessResponseInfo", string(data)}, " ")
}
