package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PacketMessage struct {

	// 16进制index
	HexIndex *string `json:"hex_index,omitempty"`

	// 16进制数列
	Hexs *[]string `json:"hexs,omitempty"`

	// utf_8字符串
	Utf8String *string `json:"utf8_String,omitempty"`
}

func (o PacketMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PacketMessage struct{}"
	}

	return strings.Join([]string{"PacketMessage", string(data)}, " ")
}
