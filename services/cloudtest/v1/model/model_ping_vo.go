package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PingVo struct {

	// ping地址
	Address *string `json:"address,omitempty"`

	// 节点名称
	SubTaskName *string `json:"sub_task_name,omitempty"`
}

func (o PingVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PingVo struct{}"
	}

	return strings.Join([]string{"PingVo", string(data)}, " ")
}
