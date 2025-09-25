package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostProtectHistoryResponseInfo struct {

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 检测时间，单位毫秒。（已废弃，请使用occur_time） **取值范围**: 最小值0，最大值4070880000000
	OccrTime *int64 `json:"occr_time,omitempty"`

	// **参数解释**: 检测时间，单位毫秒。 **取值范围**: 最小值0，最大值4070880000000
	OccurTime *int64 `json:"occur_time,omitempty"`

	// **参数解释**： 防护文件路径 **取值范围**： 字符长度0-2000位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**: 事件描述，即文件操作类型 **取值范围**: - add: 新增文件。 - delete: 删除文件。 - modify: 修改文件内容。 - attribute: 修改文件属性。 - unknown: 未知。
	FileOperation *string `json:"file_operation,omitempty"`

	// **参数解释**: 服务器IP **取值范围**: 字符长度0-64位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 进程ID **取值范围**: 最小值0，最大值2147483647
	ProcessId *int32 `json:"process_id,omitempty"`

	// **参数解释**: 进程名称 **取值范围**: 字符长度0-200位
	ProcessName *string `json:"process_name,omitempty"`

	// **参数解释**: 进程命令行 **取值范围**: 字符长度0-8192位
	ProcessCmd *string `json:"process_cmd,omitempty"`
}

func (o HostProtectHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostProtectHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"HostProtectHistoryResponseInfo", string(data)}, " ")
}
