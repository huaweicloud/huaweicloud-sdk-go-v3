package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterResult 请求成功返回的数据。
type RegisterResult struct {

	// 通道号。
	Port *string `json:"port,omitempty"`

	// 通道号类型。 - 1：普通 - 3：前缀号段 - 5：后缀号段
	PortType *int32 `json:"port_type,omitempty"`

	// 签名列表，最大长度为5。
	Sign *[]string `json:"sign,omitempty"`
}

func (o RegisterResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterResult struct{}"
	}

	return strings.Join([]string{"RegisterResult", string(data)}, " ")
}
