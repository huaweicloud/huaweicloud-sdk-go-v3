package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventProcessResInfo 进程信息
type EventProcessResInfo struct {

	// **参数解释**： 进程名称 **取值范围**： 字符长度1-128位
	ProcessName *string `json:"process_name,omitempty"`

	// **参数解释**： 进程路径 **取值范围**： 字符长度1-256位
	ProcessPath *string `json:"process_path,omitempty"`

	// **参数解释**： 进程ID **取值范围**： 最小值0，最大值2147483647
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// **参数解释**： 进程名称 **取值范围**： 最小值0，最大值2147483647
	ProcessUid *int32 `json:"process_uid,omitempty"`

	// **参数解释**： 运行进程的用户名 **取值范围**： 字符长度1-128位
	ProcessUsername *string `json:"process_username,omitempty"`

	// **参数解释**： 进程命令行 **约束限制**： 不涉及
	ProcessCmdline *string `json:"process_cmdline,omitempty"`

	// **参数解释**： 进程文件名 **取值范围**： 字符长度1-64位
	ProcessFilename *string `json:"process_filename,omitempty"`

	// **参数解释**: 进程启动时间 **取值范围**: 最小值0，最大值9223372036854775807
	ProcessStartTime *int64 `json:"process_start_time,omitempty"`

	// **参数解释**: 进程组ID **取值范围**: 最小值0，最大值2147483647
	ProcessGid *int32 `json:"process_gid,omitempty"`

	// **参数解释**: 进程有效组ID **取值范围**: 最小值0，最大值2147483647
	ProcessEgid *int32 `json:"process_egid,omitempty"`

	// **参数解释**: 进程有效用户ID **取值范围**: 最小值0，最大值9223372036854775807
	ProcessEuid *int64 `json:"process_euid,omitempty"`

	// **参数解释**： 父进程名称 **取值范围**： 字符长度1-64位
	ParentProcessName *string `json:"parent_process_name,omitempty"`

	// **参数解释**： 父进程文件路径 **取值范围**： 字符长度1-64位
	ParentProcessPath *string `json:"parent_process_path,omitempty"`

	// **参数解释**: 父进程id **取值范围**: 最小值0，最大值2147483647
	ParentProcessPid *int32 `json:"parent_process_pid,omitempty"`

	// **参数解释**: 父进程用户id **取值范围**: 最小值0，最大值2147483647
	ParentProcessUid *int32 `json:"parent_process_uid,omitempty"`

	// **参数解释**： 父进程文件命令行 **取值范围**： 字符长度1-64位
	ParentProcessCmdline *string `json:"parent_process_cmdline,omitempty"`

	// **参数解释**： 父进程文件名 **取值范围**： 字符长度1-64位
	ParentProcessFilename *string `json:"parent_process_filename,omitempty"`

	// **参数解释**: 父进程启动时间 **取值范围**: 最小值0，最大值9223372036854775807
	ParentProcessStartTime *int64 `json:"parent_process_start_time,omitempty"`

	// **参数解释**: 父进程组ID **取值范围**: 最小值0，最大值2147483647
	ParentProcessGid *int32 `json:"parent_process_gid,omitempty"`

	// **参数解释**: 父进程有效组ID **取值范围**: 最小值0，最大值2147483647
	ParentProcessEgid *int32 `json:"parent_process_egid,omitempty"`

	// **参数解释**: 父进程有效用户ID **取值范围**: 最小值0，最大值2147483647
	ParentProcessEuid *int32 `json:"parent_process_euid,omitempty"`

	// **参数解释**： 子进程名称 **取值范围**： 字符长度1-64位
	ChildProcessName *string `json:"child_process_name,omitempty"`

	// **参数解释**： 子进程文件路径 **取值范围**： 字符长度1-64位
	ChildProcessPath *string `json:"child_process_path,omitempty"`

	// **参数解释**: 子进程id **取值范围**: 最小值0，最大值2147483647
	ChildProcessPid *int32 `json:"child_process_pid,omitempty"`

	// **参数解释**: 子进程用户id **取值范围**: 最小值0，最大值2147483647
	ChildProcessUid *int32 `json:"child_process_uid,omitempty"`

	// **参数解释**： 子进程文件命令行 **取值范围**： 字符长度1-64位
	ChildProcessCmdline *string `json:"child_process_cmdline,omitempty"`

	// **参数解释**： 子进程文件名 **取值范围**： 字符长度1-64位
	ChildProcessFilename *string `json:"child_process_filename,omitempty"`

	// **参数解释**: 子进程启动时间 **取值范围**: 最小值0，最大值9223372036854775807
	ChildProcessStartTime *int64 `json:"child_process_start_time,omitempty"`

	// **参数解释**: 子进程组ID **取值范围**: 最小值0，最大值2147483647
	ChildProcessGid *int32 `json:"child_process_gid,omitempty"`

	// **参数解释**: 子进程有效组ID **取值范围**: 最小值0，最大值2147483647
	ChildProcessEgid *int32 `json:"child_process_egid,omitempty"`

	// **参数解释**: 子进程有效用户ID **取值范围**: 最小值0，最大值2147483647
	ChildProcessEuid *int32 `json:"child_process_euid,omitempty"`

	// **参数解释**： 虚拟化命令 **取值范围**： 字符长度1-64位
	VirtCmd *string `json:"virt_cmd,omitempty"`

	// **参数解释**： 虚拟化进程名称 **取值范围**： 字符长度1-64位
	VirtProcessName *string `json:"virt_process_name,omitempty"`

	// **参数解释**： 逃逸方式 **取值范围**： 字符长度1-64位
	EscapeMode *string `json:"escape_mode,omitempty"`

	// **参数解释**： 逃逸后执行的命令 **取值范围**： 字符长度1-128位
	EscapeCmd *string `json:"escape_cmd,omitempty"`

	// **参数解释**： 进程启动文件hash **取值范围**： 字符长度1-64位
	ProcessHash *string `json:"process_hash,omitempty"`

	// **参数解释**： 文件属性 **取值范围**： 字符长度1-64位
	Mode *string `json:"mode,omitempty"`

	// **参数解释**： 规则 **取值范围**： 字符长度1-64位
	Rule *int32 `json:"rule,omitempty"`

	// **参数解释**： 分数 **取值范围**： 字符长度1-64位
	Score *int32 `json:"score,omitempty"`

	// **参数解释**： 进程文件hash **取值范围**： 字符长度1-64位
	ProcessFileHash *string `json:"process_file_hash,omitempty"`

	// **参数解释**： 父进程文件hash **取值范围**： 字符长度1-64位
	ParentProcessFileHash *string `json:"parent_process_file_hash,omitempty"`

	// **参数解释**: 祖父进程id **取值范围**: 最小值1，最大值2147483647
	AncestorProcessPid *int32 `json:"ancestor_process_pid,omitempty"`

	// **参数解释**： 祖父进程命令行 **取值范围**： 字符长度1-64位
	AncestorProcessCmdline *string `json:"ancestor_process_cmdline,omitempty"`

	// **参数解释**： 祖父进程路径 **取值范围**： 字符长度1-64位
	AncestorProcessPath *string `json:"ancestor_process_path,omitempty"`

	// **参数解释**: 操作类型 **取值范围**: 最小值1，最大值2147483647
	OperateType *int32 `json:"operate_type,omitempty"`

	// **参数解释**: 会话id **取值范围**: 最小值1，最大值2147483647
	SessionId *int32 `json:"session_id,omitempty"`
}

func (o EventProcessResInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventProcessResInfo struct{}"
	}

	return strings.Join([]string{"EventProcessResInfo", string(data)}, " ")
}
