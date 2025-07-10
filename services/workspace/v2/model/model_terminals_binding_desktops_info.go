package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TerminalsBindingDesktopsInfo struct {

	// MAC绑定策略ID。
	Id *string `json:"id,omitempty"`

	// 终端MAC地址。
	Mac *string `json:"mac,omitempty"`

	// 虚拟机名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o TerminalsBindingDesktopsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TerminalsBindingDesktopsInfo struct{}"
	}

	return strings.Join([]string{"TerminalsBindingDesktopsInfo", string(data)}, " ")
}
