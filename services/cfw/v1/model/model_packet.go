package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Packet struct {

	// 十六进制码
	HexIndex *string `json:"hex_index,omitempty"`

	// utf8字符串
	Utf8String *string `json:"utf8_string,omitempty"`

	// 十六进制单个字节码数组
	Hexs *[]string `json:"hexs,omitempty"`
}

func (o Packet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Packet struct{}"
	}

	return strings.Join([]string{"Packet", string(data)}, " ")
}
