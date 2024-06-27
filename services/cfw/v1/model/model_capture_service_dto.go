package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaptureServiceDto struct {

	// 目的端口
	DestPort *string `json:"dest_port,omitempty"`

	// 协议类型:TCP为6，UDP为17，ICMP为1，ICMPV6为58，ANY为-1，手动类型不为空，自动类型为空
	Protocol int32 `json:"protocol"`

	// 源端口
	SourcePort *string `json:"source_port,omitempty"`
}

func (o CaptureServiceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaptureServiceDto struct{}"
	}

	return strings.Join([]string{"CaptureServiceDto", string(data)}, " ")
}
