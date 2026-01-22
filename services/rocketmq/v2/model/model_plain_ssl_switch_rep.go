package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PlainSslSwitchRep struct {

	// **参数解释**： 协议模式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TlsMode *string `json:"tls_mode,omitempty"`
}

func (o PlainSslSwitchRep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlainSslSwitchRep struct{}"
	}

	return strings.Join([]string{"PlainSslSwitchRep", string(data)}, " ")
}
