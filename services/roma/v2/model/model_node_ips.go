package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点内部ip信息
type NodeIps struct {

	// 自定义后端服务livedta组件节点ip列表
	Livedata *[]string `json:"livedata,omitempty" xml:"livedata"`

	// API网关shubao组件节点ip列表
	Shubao *[]string `json:"shubao,omitempty" xml:"shubao"`
}

func (o NodeIps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeIps struct{}"
	}

	return strings.Join([]string{"NodeIps", string(data)}, " ")
}
