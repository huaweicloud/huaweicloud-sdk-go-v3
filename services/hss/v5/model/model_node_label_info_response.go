package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeLabelInfoResponse struct {

	// **参数解释**: 节点标签名称。 **约束限制**: 不涉及 **取值范围**: 字符长度0-512位 **默认取值**: 不涉及
	Label *string `json:"label,omitempty"`

	// **参数解释**: 节点名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	Children *[]string `json:"children,omitempty"`
}

func (o NodeLabelInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeLabelInfoResponse struct{}"
	}

	return strings.Join([]string{"NodeLabelInfoResponse", string(data)}, " ")
}
