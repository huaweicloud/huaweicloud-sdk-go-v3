package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点类型对象
type NodeTypes struct {
	// 节点类型详细

	Detail []Detail `json:"detail"`
	// 节点类型ID

	Id string `json:"id"`

	SpecName string `json:"spec_name"`
}

func (o NodeTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypes struct{}"
	}

	return strings.Join([]string{"NodeTypes", string(data)}, " ")
}
