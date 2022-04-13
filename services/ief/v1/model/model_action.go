package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点
type Action struct {
	// 启用，停用边缘节点，支持start/stop

	Action string `json:"action"`
}

func (o Action) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Action struct{}"
	}

	return strings.Join([]string{"Action", string(data)}, " ")
}
