package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 要加入的节点详情
type EdgeNodeAdded struct {
	// 终端设备和节点关系的名称，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64

	Relation *string `json:"relation,omitempty"`
	// 终端设备和节点关系的描述，最大长度64，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\

	Comment *string `json:"comment,omitempty"`
	// 节点ID列表，一个设备只可以被绑定于一个边缘节点。

	NodeIds []string `json:"node_ids"`
}

func (o EdgeNodeAdded) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeAdded struct{}"
	}

	return strings.Join([]string{"EdgeNodeAdded", string(data)}, " ")
}
