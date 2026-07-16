package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerPhyInfo struct {

	// **参数解释**：Lite Server实例资源ID。 **取值范围**：长度为[8,36]个字符。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**：Tor信息。 **取值范围**：多个ip信息，IPv4格式。
	NetworkNodes *[]string `json:"network_nodes,omitempty"`
}

func (o ServerPhyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerPhyInfo struct{}"
	}

	return strings.Join([]string{"ServerPhyInfo", string(data)}, " ")
}
