package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtocolReqVo struct {
	Basic *BasicInfoVo `json:"basic,omitempty"`

	// 请求头
	Headers map[string]string `json:"headers,omitempty"`

	// 方法
	Method *string `json:"method,omitempty"`

	// 请求body体
	RequestBody *string `json:"request_body,omitempty"`

	// url
	Url *string `json:"url,omitempty"`
}

func (o ProtocolReqVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtocolReqVo struct{}"
	}

	return strings.Join([]string{"ProtocolReqVo", string(data)}, " ")
}
