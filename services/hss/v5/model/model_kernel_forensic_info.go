package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelForensicInfo 内核取证信息
type KernelForensicInfo struct {

	// **参数解释**： 地址 **取值范围**： 不涉及
	ReadAddr *int32 `json:"read_addr,omitempty"`

	// **参数解释**： 系统调用编号 **取值范围**： 不涉及
	SyscallNum *int32 `json:"syscall_num,omitempty"`

	// **参数解释**： 系统函数 **取值范围**： 不涉及
	Function *string `json:"function,omitempty"`

	// **参数解释**： 内核模块 **取值范围**： 不涉及
	Module *string `json:"module,omitempty"`

	// **参数解释**： 扩展信息 **取值范围**： 不涉及
	ExtInfo *string `json:"ext_info,omitempty"`
}

func (o KernelForensicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelForensicInfo struct{}"
	}

	return strings.Join([]string{"KernelForensicInfo", string(data)}, " ")
}
