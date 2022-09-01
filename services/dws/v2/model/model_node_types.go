package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点类型对象
type NodeTypes struct {

	// 节点类型名称。
	SpecName string `json:"spec_name" xml:"spec_name"`

	// 节点类型详细。
	Detail []Detail `json:"detail" xml:"detail"`

	// 节点类型ID。
	Id string `json:"id" xml:"id"`
}

func (o NodeTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypes struct{}"
	}

	return strings.Join([]string{"NodeTypes", string(data)}, " ")
}
