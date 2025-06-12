package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PlainSslSwitchRep struct {

	// 协议模式。
	TlsMode *string `json:"tls_mode,omitempty"`
}

func (o PlainSslSwitchRep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlainSslSwitchRep struct{}"
	}

	return strings.Join([]string{"PlainSslSwitchRep", string(data)}, " ")
}
