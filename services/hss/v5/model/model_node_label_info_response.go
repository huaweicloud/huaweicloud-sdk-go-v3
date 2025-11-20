package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeLabelInfoResponse struct {

	// **参数解释**: 节点标签名称。 **取值范围**: 字符长度1-512位
	Label *string `json:"label,omitempty"`

	// **参数解释**: 节点名称 **取值范围**: 字符长度0-256位
	Children *[]string `json:"children,omitempty"`
}

func (o NodeLabelInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeLabelInfoResponse struct{}"
	}

	return strings.Join([]string{"NodeLabelInfoResponse", string(data)}, " ")
}
