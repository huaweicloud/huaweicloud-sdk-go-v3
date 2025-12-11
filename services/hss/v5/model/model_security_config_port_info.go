package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityConfigPortInfo 危险端口信息
type SecurityConfigPortInfo struct {

	// **参数解释**： 端口号 **取值范围**： 0-65535
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 类型 **取值范围**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 端口危险程度 **取值范围**： - normal：正常端口 - danger：危险端口 - unknown：未知端口
	Status *string `json:"status,omitempty"`

	// **参数解释**： 端口状态 **取值范围**： - 0：未处理 - 1：已忽略 - 2：无需处理
	PortStatus *int32 `json:"port_status,omitempty"`

	// **参数解释**： 端口描述 **取值范围**： 不涉及
	PortDesc *string `json:"port_desc,omitempty"`
}

func (o SecurityConfigPortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityConfigPortInfo struct{}"
	}

	return strings.Join([]string{"SecurityConfigPortInfo", string(data)}, " ")
}
