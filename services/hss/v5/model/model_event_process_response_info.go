package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventProcessResponseInfo 进程信息
type EventProcessResponseInfo struct {

	// **参数解释**： 进程名称 **取值范围**： 字符长度1-256位
	ProcessName *string `json:"process_name,omitempty"`

	// **参数解释**： 进程文件路径 **取值范围**： 字符长度1-256位
	ProcessPath *string `json:"process_path,omitempty"`

	// **参数解释**： 进程ID **取值范围**： 最小值0，最大值2147483647
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// **参数解释**： 进程用户ID **取值范围**： 最小值0，最大值2147483647
	ProcessUid *int32 `json:"process_uid,omitempty"`

	// **参数解释**： 运行进程的用户名 **取值范围**： 字符长度1-256位
	ProcessUsername *string `json:"process_username,omitempty"`

	// **参数解释**： 进程文件命令行 **取值范围**： 字符长度1-256位
	ProcessCmdline *string `json:"process_cmdline,omitempty"`

	// **参数解释**： 进程文件名 **取值范围**： 字符长度1-256位
	ProcessFilename *string `json:"process_filename,omitempty"`

	// **参数解释**： 进程启动时间 **取值范围**： 最小值0，最大值9223372036854775807
	ProcessStartTime *int64 `json:"process_start_time,omitempty"`

	// **参数解释**： 进程组ID **取值范围**： 最小值0，最大值2147483647
	ProcessGid *int32 `json:"process_gid,omitempty"`

	// **参数解释**： 进程有效组ID **取值范围**： 最小值0，最大值2147483647
	ProcessEgid *int32 `json:"process_egid,omitempty"`

	// **参数解释**： 进程有效用户ID **取值范围**： 最小值0，最大值2147483647
	ProcessEuid *int32 `json:"process_euid,omitempty"`

	// **参数解释**： 祖父进程文件路径 **取值范围**： 字符长度1-256位
	AncestorProcessPath *string `json:"ancestor_process_path,omitempty"`

	// **参数解释**： 祖父进程ID **取值范围**： 最小值0，最大值2147483647
	AncestorProcessPid *int32 `json:"ancestor_process_pid,omitempty"`

	// **参数解释**： 祖父进程文件命令行 **取值范围**： 字符长度1-512位
	AncestorProcessCmdline *string `json:"ancestor_process_cmdline,omitempty"`

	// **参数解释**： 父进程名称 **取值范围**： 字符长度1-256位
	ParentProcessName *string `json:"parent_process_name,omitempty"`

	// **参数解释**： 父进程文件路径 **取值范围**： 字符长度1-256位
	ParentProcessPath *string `json:"parent_process_path,omitempty"`

	// **参数解释**： 父进程ID **取值范围**： 最小值0，最大值2147483647
	ParentProcessPid *int32 `json:"parent_process_pid,omitempty"`

	// **参数解释**： 父进程用户ID **取值范围**： 最小值0，最大值2147483647
	ParentProcessUid *int32 `json:"parent_process_uid,omitempty"`

	// **参数解释**： 父进程文件命令行 **取值范围**： 字符长度1-512位
	ParentProcessCmdline *string `json:"parent_process_cmdline,omitempty"`

	// **参数解释**： 父进程文件名 **取值范围**： 字符长度1-256位
	ParentProcessFilename *string `json:"parent_process_filename,omitempty"`

	// **参数解释**： 父进程启动时间 **取值范围**： 最小值0，最大值9223372036854775807
	ParentProcessStartTime *int64 `json:"parent_process_start_time,omitempty"`

	// **参数解释**： 父进程组ID **取值范围**： 最小值0，最大值2147483647
	ParentProcessGid *int32 `json:"parent_process_gid,omitempty"`

	// **参数解释**： 父进程有效组ID **取值范围**： 最小值0，最大值2147483647
	ParentProcessEgid *int32 `json:"parent_process_egid,omitempty"`

	// **参数解释**： 父进程有效用户ID **取值范围**： 最小值0，最大值2147483647
	ParentProcessEuid *int32 `json:"parent_process_euid,omitempty"`

	// **参数解释**： 子进程名称 **取值范围**： 字符长度1-256位
	ChildProcessName *string `json:"child_process_name,omitempty"`

	// **参数解释**： 子进程文件路径 **取值范围**： 字符长度1-256位
	ChildProcessPath *string `json:"child_process_path,omitempty"`

	// **参数解释**： 子进程id **取值范围**： 最小值0，最大值2147483647
	ChildProcessPid *int32 `json:"child_process_pid,omitempty"`

	// **参数解释**： 子进程用户id **取值范围**： 最小值0，最大值2147483647
	ChildProcessUid *int32 `json:"child_process_uid,omitempty"`

	// **参数解释**： 子进程文件命令行 **取值范围**： 字符长度1-256位
	ChildProcessCmdline *string `json:"child_process_cmdline,omitempty"`

	// **参数解释**： 子进程文件名 **取值范围**： 字符长度1-256位
	ChildProcessFilename *string `json:"child_process_filename,omitempty"`

	// **参数解释**： 子进程启动时间 **取值范围**： 最小值0，最大值9223372036854775807
	ChildProcessStartTime *int64 `json:"child_process_start_time,omitempty"`

	// **参数解释**： 子进程组ID **取值范围**： 最小值0，最大值2147483647
	ChildProcessGid *int32 `json:"child_process_gid,omitempty"`

	// **参数解释**： 子进程有效组ID **取值范围**： 最小值0，最大值2147483647
	ChildProcessEgid *int32 `json:"child_process_egid,omitempty"`

	// **参数解释**： 子进程有效用户ID **取值范围**： 最小值0，最大值2147483647
	ChildProcessEuid *int32 `json:"child_process_euid,omitempty"`

	// **参数解释**： 虚拟化命令 **取值范围**： 字符长度1-256位
	VirtCmd *string `json:"virt_cmd,omitempty"`

	// **参数解释**： 虚拟化进程名称 **取值范围**： 字符长度1-256位
	VirtProcessName *string `json:"virt_process_name,omitempty"`

	// **参数解释**： 逃逸方式 **取值范围**： 字符长度1-256位
	EscapeMode *string `json:"escape_mode,omitempty"`

	// **参数解释**： 逃逸后后执行的命令 **取值范围**： 字符长度1-256位
	EscapeCmd *string `json:"escape_cmd,omitempty"`

	// **参数解释**： 进程启动文件hash **取值范围**： 字符长度1-256位
	ProcessHash *string `json:"process_hash,omitempty"`

	// **参数解释**： 进程文件hash **取值范围**： 字符长度1-256位
	ProcessFileHash *string `json:"process_file_hash,omitempty"`

	// **参数解释**： 父进程文件hash **取值范围**： 字符长度1-256位
	ParentProcessFileHash *string `json:"parent_process_file_hash,omitempty"`

	// 是否阻断成功，1阻断成功 0阻断失败
	Block *int32 `json:"block,omitempty"`
}

func (o EventProcessResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventProcessResponseInfo struct{}"
	}

	return strings.Join([]string{"EventProcessResponseInfo", string(data)}, " ")
}
