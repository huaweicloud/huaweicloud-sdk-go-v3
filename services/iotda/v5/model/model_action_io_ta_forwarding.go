package model

import (
	"encoding/json"

	"strings"
)

// 转发IoTA服务消息结构
type ActionIoTaForwarding struct {
	// IoTA服务对应的region区域
	RegionName string `json:"region_name"`
	// IoTA服务对应的projectId信息
	ProjectId string `json:"project_id"`
}

func (o ActionIoTaForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ActionIoTaForwarding struct{}"
	}

	return strings.Join([]string{"ActionIoTaForwarding", string(data)}, " ")
}
