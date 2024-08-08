package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ForwardingInfo 企业版实例的SNAT配置信息。
type ForwardingInfo struct {

	// **参数说明**：NAT网关绑定的EIP
	Eip *string `json:"eip,omitempty"`

	// **参数说明**：是否启用SNAT配置 **取值范围**： - true: SNAT配置已启用 - false: SNAT配置未启用
	EnableSnat *bool `json:"enable_snat,omitempty"`
}

func (o ForwardingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForwardingInfo struct{}"
	}

	return strings.Join([]string{"ForwardingInfo", string(data)}, " ")
}
