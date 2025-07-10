package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTerminalsBindingDesktopsRequest Request Object
type ListTerminalsBindingDesktopsRequest struct {

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// mac地址。
	Mac *string `json:"mac,omitempty"`

	// 起始数。
	Offset int32 `json:"offset"`

	// 数量。
	Limit int32 `json:"limit"`

	// 是否只查询结果总条数。
	CountOnly *bool `json:"count_only,omitempty"`
}

func (o ListTerminalsBindingDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTerminalsBindingDesktopsRequest struct{}"
	}

	return strings.Join([]string{"ListTerminalsBindingDesktopsRequest", string(data)}, " ")
}
