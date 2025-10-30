package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessForensicInfo 进程链取证信息
type ProcessForensicInfo struct {

	// **参数解释**： 进程名称 **取值范围**： 不涉及
	ProcessName *string `json:"process_name,omitempty"`

	// **参数解释**： 进程文件路径 **取值范围**： 不涉及
	ProcessPath *string `json:"process_path,omitempty"`

	// **参数解释**： 进程id **取值范围**： 不涉及
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// **参数解释**： 进程用户id **取值范围**： 不涉及
	ProcessUid *int32 `json:"process_uid,omitempty"`

	// **参数解释**： 进程组id **取值范围**： 不涉及
	ProcessGid *int32 `json:"process_gid,omitempty"`

	// **参数解释**： 进程有效组id **取值范围**： 不涉及
	ProcessEgid *int32 `json:"process_egid,omitempty"`

	// **参数解释**： 进程有效用户id **取值范围**： 不涉及
	ProcessEuid *int32 `json:"process_euid,omitempty"`

	// **参数解释**： 运行进程的用户名 **取值范围**： 不涉及
	ProcessUsername *string `json:"process_username,omitempty"`

	// **参数解释**： 进程文件命令行 **取值范围**： 不涉及
	ProcessCmdline *string `json:"process_cmdline,omitempty"`

	// **参数解释**： 进程启动时间 **取值范围**： 不涉及
	ProcessStartTime *int64 `json:"process_start_time,omitempty"`

	// **参数解释**： 进程文件hash **取值范围**： 不涉及
	ProcessFileHash *string `json:"process_file_hash,omitempty"`

	// **参数解释**： 祖父进程id **取值范围**： 不涉及
	AncestorProcessPid *int32 `json:"ancestor_process_pid,omitempty"`

	// **参数解释**： 祖父进程命令行 **取值范围**： 不涉及
	AncestorProcessCmdline *string `json:"ancestor_process_cmdline,omitempty"`

	// **参数解释**： 祖父进程路径 **取值范围**： 不涉及
	AncestorProcessPath *string `json:"ancestor_process_path,omitempty"`

	// **参数解释**： 会话id **取值范围**： 不涉及
	SessionId *int32 `json:"session_id,omitempty"`

	// **参数解释**： 威胁事件数 **取值范围**： 不涉及
	EventNum *int32 `json:"event_num,omitempty"`

	// **参数解释**： 节点类型 **取值范围**： - 0：进程 - 1：注册表 - 2：文件
	Type *string `json:"type,omitempty"`
}

func (o ProcessForensicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessForensicInfo struct{}"
	}

	return strings.Join([]string{"ProcessForensicInfo", string(data)}, " ")
}
