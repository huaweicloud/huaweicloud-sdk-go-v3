package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTerminalsBindingDesktopsInfo struct {

	// 行号,用于批量导入。
	Line *int32 `json:"line,omitempty"`

	// 终端mac地址。
	Mac *string `json:"mac,omitempty"`

	// 桌面名称，用于批量导入。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o CreateTerminalsBindingDesktopsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTerminalsBindingDesktopsInfo struct{}"
	}

	return strings.Join([]string{"CreateTerminalsBindingDesktopsInfo", string(data)}, " ")
}
