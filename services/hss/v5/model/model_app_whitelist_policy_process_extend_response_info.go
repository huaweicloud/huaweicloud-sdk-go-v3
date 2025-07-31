package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppWhitelistPolicyProcessExtendResponseInfo 策略识别进程信息
type AppWhitelistPolicyProcessExtendResponseInfo struct {

	// **参数解释**： 进程名称 **取值范围**： 字符长度1-128位
	ProcessName *string `json:"process_name,omitempty"`

	// **参数解释**： 进程路径 **取值范围**： 字符长度1-256位
	ProcessPath *string `json:"process_path,omitempty"`

	// 进程hash
	ProcessHash *string `json:"process_hash,omitempty"`

	// **参数解释**: 容器ID **取值范围**: 字符长度1-128位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 进程命令行 **取值范围**: 不涉及
	Cmdline *string `json:"cmdline,omitempty"`

	// **参数解释**: 文件大小 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	FileSize *int64 `json:"file_size,omitempty"`
}

func (o AppWhitelistPolicyProcessExtendResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppWhitelistPolicyProcessExtendResponseInfo struct{}"
	}

	return strings.Join([]string{"AppWhitelistPolicyProcessExtendResponseInfo", string(data)}, " ")
}
