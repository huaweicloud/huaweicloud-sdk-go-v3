package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeIps struct {

	// livedata节点IP地址列表
	Livedata *[]string `json:"livedata,omitempty"`

	// shubao节点IP地址列表
	Shubao *[]string `json:"shubao,omitempty"`
}

func (o NodeIps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeIps struct{}"
	}

	return strings.Join([]string{"NodeIps", string(data)}, " ")
}
